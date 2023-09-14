import { IFileConfigContentSummary } from './config';
import { IVariableEditParams } from './variable';

// 版本下的脚本配置
export interface IDiffDetail {
  contentType: 'file'|'text';
  base: {
    content: string|IFileConfigContentSummary;
    language?: string;
    variables?: IVariableEditParams[];
  },
  current: {
    content: string|IFileConfigContentSummary;
    language?: string;
    variables?: IVariableEditParams[];
  }
}
