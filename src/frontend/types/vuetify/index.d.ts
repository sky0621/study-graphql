export interface DataTableHeader {
  text: string
  value: string
  sortable: boolean
  width?: number
  filter?: (value: any, search: string, item: any) => boolean
}
