import type { ComponentInternalInstance } from 'vue'

export interface DialogOptions {
  title?: string
}

export interface DialogShowFunc {
  (message: string, options?: DialogOptions = {}): Promise<void>
}

export interface WarnDialogInstance extends ComponentInternalInstance {
  callback(): void
}

export interface ConfirmDialogInstance extends ComponentInternalInstance {
  closeCallback(): void
  confirmCallback(): void
}
