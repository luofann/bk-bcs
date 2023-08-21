<script lang="ts" setup>
  import { computed, ref, watch } from 'vue';
  import SHA256 from 'crypto-js/sha256'
  import WordArray from 'crypto-js/lib-typedarrays'
  import { Message } from 'bkui-vue';
  import { Done, TextFill } from 'bkui-vue/lib/icon'
  import { ITemplateVersionEditingData } from '../../../../../../types/template'
  import { IFileConfigContentSummary } from '../../../../../../types/config'
  import { stringLengthInBytes } from '../../../../../utils/index'
  import { transFileToObject, fileDownload } from '../../../../../utils/file'
  import { updateTemplateContent, downloadTemplateContent, createTemplateVersion } from '../../../../../api/template'
  import CodeEditor from '../../../../../components/code-editor/index.vue'
  import PermissionInputPicker from '../../../../../components/permission-input-picker.vue';
  import CreateVersionConfirmDialog from './create-version-confirm-dialog.vue';


  const props = defineProps<{
    spaceId: string;
    templateSpaceId: number;
    templateId: number;
    versionId: number;
    versionName: string;
    type: string;
    data: ITemplateVersionEditingData;
  }>()

  const emits = defineEmits(['created', 'close'])

  const formData = ref<ITemplateVersionEditingData>({
    revision_memo: '',
    file_type: '',
    file_mode: '',
    user: '',
    user_group: '',
    privilege: '',
    sign: '',
    byte_size: 0
  })
  const formRef = ref()
  const stringContent = ref('')
  const fileContent = ref<IFileConfigContentSummary|File>()
  const isFileChanged = ref(false)
  const uploadPending = ref(false)
  const submitPending = ref(false)
  const isConfirmDialogShow = ref(false)

  const isViewMode = computed(() => {
    return props.type === 'view'
  })

  // 传入到bk-upload组件的文件对象
  const fileList = computed(() => {
    return fileContent.value ? [transFileToObject(<File>fileContent.value)] : []
  })

  watch(() => props.data, val => {
    formData.value = { ...val }
  }, { immediate: true })

  const handleFileUpload = (option: { file: File }) => {
    isFileChanged.value = true
    return new Promise(resolve => {
      fileContent.value = option.file
      uploadContent().then(res => {
        resolve(res)
      })
    })
  }

  // 上传配置内容
  const uploadContent =  async () => {
    const signature = await getSignature()
    const data = formData.value.file_type === 'binary' ? fileContent.value : stringContent.value
    uploadPending.value = true
    // @ts-ignore
    return updateTemplateContent(props.spaceId, props.templateSpaceId, data, signature).then(() => {
      if (formData.value.file_type === 'binary') {
        formData.value.byte_size = Number((<IFileConfigContentSummary|File>fileContent.value).size)
      } else {
        formData.value.byte_size = new Blob([stringContent.value]).size
      }
      formData.value.sign = signature
      uploadPending.value = false
    })
  }

  // 生成文件或文本的sha256
  const getSignature = async () => {
    if (props.data.file_type === 'binary') {
      if (isFileChanged.value) {
        return new Promise(resolve => {
          const reader = new FileReader()
          // @ts-ignore
          reader.readAsArrayBuffer(fileContent.value)
          reader.onload = () => {
            const wordArray = WordArray.create(reader.result);
            resolve(SHA256(wordArray).toString())
          }
        })
      }
      return (fileContent.value as IFileConfigContentSummary).signature
    }
    return SHA256(stringContent.value).toString()
  }

    // 下载已上传文件
    const handleDownloadFile = async () => {
    const { signature, name } = <IFileConfigContentSummary>fileContent.value
    const res = await downloadTemplateContent(props.spaceId, props.templateSpaceId, signature)
    fileDownload(res, `${name}.bin`)
  }

  const validate = async() => {
    await formRef.value.validate()
    if (formData.value.file_type === 'binary'){
      if (fileList.value.length === 0) {
        Message({ theme: 'error', message: '请上传文件' })
        return false
      }
    } else if (formData.value.file_type === 'text') {
      if (stringLengthInBytes(stringContent.value) > 1024 * 1024 * 50 ) {
        Message({ theme: 'error', message: '配置内容不能超过50M' })
        return false
      }
    }
    return true
  }

  const handleSubmitClick = async() => {
    const result = await validate()
    if (!result) return
    isConfirmDialogShow.value = true
  }

  const triggerCreate = async() => {
    try {
      submitPending.value = true
      if (formData.value.file_type !== 'binary') {
        await uploadContent()
      }
      const res = await createTemplateVersion(props.spaceId, props.templateSpaceId, props.templateId, formData.value)
      emits('created', res.id)
      Message({
        theme: 'success',
        message: '创建版本成功'
      })
    } catch (e) {
      console.log(e)
    } finally {
      submitPending.value = false
    }
  }

</script>
<template>
  <div class="version-editor">
    <div class="header-wrapper">
      <div class="title">
        <div v-if="props.type === 'create'" class="create-version-title">新建版本</div>
        <div v-else class="version-view-title">
          {{ props.versionName }}
          <span class="cited-info">被引用：{{ 0 }}</span>
        </div>
      </div>
    </div>
    <div :class="['template-config-content-wrapper', { 'view-mode': isViewMode }]">
      <div v-if="!isViewMode" class="config-form">
        <bk-form ref="formRef" form-type="vertical">
          <bk-form-item label="版本描述">
            <bk-input v-model="formData.revision_memo" type="textarea" :rows="4" :maxlength="200" />
          </bk-form-item>
          <bk-form-item label="文件权限" required>
            <PermissionInputPicker v-model="formData.privilege" />
          </bk-form-item>
          <bk-form-item label="用户" required>
            <bk-input v-model="formData.user" />
          </bk-form-item>
          <bk-form-item label="用户组" required>
            <bk-input v-model="formData.user_group" />
          </bk-form-item>
        </bk-form>
      </div>
      <div class="config-content">
        <div v-if="props.data.file_type === 'binary'" class="file-uploader-wrapper">
          <bk-upload
            class="config-uploader"
            url=""
            theme="button"
            tip="支持扩展名：.bin，文件大小100M以内"
            :size="100"
            :disabled="isViewMode"
            :multiple="false"
            :files="fileList"
            :custom-request="handleFileUpload">
            <template #file="{ file }">
              <div class="file-wrapper">
                <Done class="done-icon"/>
                <TextFill class="file-icon" />
                <div v-bk-ellipsis class="name" @click="handleDownloadFile">{{ file.name }}</div>
                ({{ file.size }})
              </div>
            </template>
          </bk-upload>
        </div>
        <CodeEditor v-else v-model="stringContent" :editable="!isViewMode" />
      </div>
    </div>
    <div class="action-btns">
      <bk-button class="submit-btn" theme="primary" @click="handleSubmitClick">提交</bk-button>
      <bk-button class="cancel-btn" @click="emits('close')">取消</bk-button>
    </div>
  </div>
  <CreateVersionConfirmDialog
    v-model:show="isConfirmDialogShow"
    :spaceId="props.spaceId"
    :template-space-id="props.templateSpaceId"
    :templateId="props.templateId"
    :version-id="props.versionId"
    :pending="submitPending"
    @confirm="triggerCreate"/>
</template>
<style lang="scss" scoped>
  .version-editor {
    height: 100%;
  }
  .header-wrapper {
    height: 40px;
    background: #242424;
    box-shadow: 0 2px 4px 0 #00000029;
  }
  .title {
    display: flex;
    align-items: center;
    padding: 0 24px;
    height: 100%;
    .create-version-title {
      line-height: 20px;
      font-size: 14px;
      color: #8a8f99;
    }
    .version-view-title {
      line-height: 20px;
      font-size: 14px;
      color: #c4c6cc;
      .cited-info {
        margin-left: 16px;
        padding-left: 16px;
        color: #63656e;
        border-left: 1px solid #63656e;
      }
    }
  }
  .template-config-content-wrapper {
    display: flex;
    align-items: flex-start;
    height: calc(100% - 86px);
    &.view-mode {
      height: calc(100% - 40px);
      .config-content {
        width: 100%;
      }
    }
    .config-form {
      padding: 24px;
      width: 260px;
      height: 100%;
      background: #2a2a2a;
      overflow: auto;
      :deep(.bk-form) {
        .bk-form-label {
          font-size: 12px;
          color: #979ba5;
        }
        .bk-input {
          border: 1px solid #63656e;
        }
        .bk-input--text {
          background: transparent;
          color: #c4c6cc;
          &::placeholder {
            color: #63656e;
          }
        }
        .bk-textarea{
          background: transparent;
          border: 1px solid #63656e;
          textarea {
            color: #c4c6cc;
            background: transparent;
            &::placeholder {
              color: #63656e;
            }
          }
        }
      }
    }
    .config-content {
      width: calc(100% - 260px);
      height: 100%;
    }
  }
  .permission-input-picker {
    :deep(.perm-panel-trigger) {
      background: #1e3250;
    }
  }
  .file-uploader-wrapper {
    padding: 24px;
    height: 100%;
  }
  .config-uploader {
    :deep(.bk-upload-list__item) {
      padding: 0;
      border: none;
    }
    :deep(.bk-upload-list--disabled .bk-upload-list__item) {
      pointer-events: inherit;
    }
    .file-wrapper {
      display: flex;
      align-items: center;
      color: #979ba5;
      font-size: 12px;
      .done-icon {
        font-size: 20px;
        color: #2dcb56;
      }
      .file-icon {
        margin: 0 6px 0 0;
      }
      .name {
        max-width: 360px;
        margin-right: 4px;
        color: #63656e;
        white-space: nowrap;
        text-overflow: ellipsis;
        overflow: hidden;
        cursor: pointer;
        &:hover {
          color: #3a84ff;
          text-decoration: underline;
        }
      }
    }
  }
  .action-btns {
    padding: 7px 24px;
    background: #2a2a2a;
    box-shadow: 0 -1px 0 0 #141414;
    .submit-btn {
      margin-right: 8px;
      min-width: 120px;
    }
    .cancel-btn {
      min-width: 88px;
      background: transparent;
      border-color: #979ba5;
      color: #979ba5;
    }
  }
</style>