<template>
  <div class="group-select-wrapper">
    <h3 class="title">{{ t('选择上线实例范围') }}</h3>
    <div class="select-group-radius">
      <bk-radio-group :model-value="type" @change="handleTypeChange">
        <bk-radio label="all">
          {{ t('全部实例上线') }}
        </bk-radio>
        <bk-radio v-if="!props.disableSelect" label="select">
          {{ t('选择分组实例上线') }}
          <GroupTree
            v-if="type === 'select'"
            :group-list="props.groupList"
            :version-list="props.versionList"
            :released-groups="props.releasedGroups"
            :value="props.value"
            @change="handleSelectGroup">
          </GroupTree>
        </bk-radio>
      </bk-radio-group>
    </div>
  </div>
</template>
<script setup lang="ts">
  import { ref } from 'vue';
  import { useI18n } from 'vue-i18n';
  import { IGroupToPublish } from '../../../../../../../../types/group';
  import { IConfigVersion } from '../../../../../../../../types/config';
  import GroupTree from './group-tree.vue';

  const { t } = useI18n();
  const props = withDefaults(
    defineProps<{
      groupList: IGroupToPublish[];
      versionList: IConfigVersion[];
      releaseType: string;
      releasedGroups?: number[];
      disableSelect: boolean;
      value: IGroupToPublish[];
    }>(),
    {
      groupList: () => [],
      versionList: () => [],
      releasedGroups: () => [],
      disableSelect: false,
      value: () => [],
    },
  );

  const emits = defineEmits(['releaseTypeChange', 'change']);

  const type = ref(props.releaseType);

  // 切换选择分组类型
  const handleTypeChange = (val: string) => {
    type.value = val;
    if (val === 'all') {
      handleSelectGroup(props.groupList);
    } else if (val === 'select') {
      const list = props.groupList.filter((group) => props.releasedGroups.includes(group.id));
      handleSelectGroup(list);
    }
    emits('releaseTypeChange', val);
  };

  const handleSelectGroup = (val: IGroupToPublish[]) => {
    emits('change', val);
  };
</script>
<style lang="scss" scoped>
  .group-select-wrapper {
    height: 100%;
  }
  .title {
    margin: 0 0 22px;
    line-height: 19px;
    font-size: 14px;
    font-weight: 700;
    color: #63656e;
  }
  .bk-radio-group {
    display: block;
  }
  .bk-radio {
    display: block;
    margin: 0 0 12px;
    &:last-of-type {
      margin-bottom: 0;
    }
    :deep(.bk-radio-label) {
      font-size: 12px;
    }
  }
</style>
