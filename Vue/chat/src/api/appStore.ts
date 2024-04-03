import { get, post } from '@/utils/request'

/* 查询app分组 */
export function fetchQueryAppCatsAPI<T>(): Promise<T> {
  return get<T>({ url: '/app/queryCats' })
}

/*  查询全量app列表 */
export function fetchQueryAppsAPI<T>(data?: { page: number, size: number }): Promise<T> {
  return get<T>({
    url: '/app/list',
		data
  })
}

/*  查询收藏的app  */
export function fetchQueryMineAppsAPI<T>(): Promise<T> {
  return get<T>({
    url: '/app/mineApps',
  })
}

/* 收藏app */
export function fetchCollectAppAPI<T>(data: { appId: number }): Promise<T> {
  return post<T>({ url: '/app/collect', data })
}

/* 收藏app */
export function fetchCustomAppAPI<T>(data: any): Promise<T> {
  return post<T>({ url: '/app/customApp', data })
}

/* 删除app */
export function fetchDelMineAppAPI<T>(data: any): Promise<T> {
  return post<T>({ url: '/app/delMineApp', data })
}
