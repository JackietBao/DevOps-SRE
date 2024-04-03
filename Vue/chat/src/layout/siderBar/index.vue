<script setup lang='ts'>
import { computed, defineAsyncComponent, ref,  watch, onMounted, nextTick } from 'vue'
import { NTooltip, useMessage, NScrollbar } from 'naive-ui'

import { useRoute, useRouter } from 'vue-router'
import { SvgIcon } from '@/components/common'
import {fetchQueryMenuAPI} from '@/api/config'
import { useAppStore, useAuthStore, useGlobalStoreWithOut } from '@/store'
import { useBasicLayout } from '@/hooks/useBasicLayout'
const Setting = defineAsyncComponent(() => import('@/components/common/Setting/index.vue'))

const appStore = useAppStore()
const authStore = useAuthStore()
const useGlobalStore = useGlobalStoreWithOut()
const message = useMessage()
const track = ref(null)
appStore.setEnv()

const avatar = computed(() => authStore.userInfo.avatar)
const route = useRoute()
const router = useRouter()
const show = ref(false)
const isLogin = computed(() => authStore.isLogin)
const darkMode = computed(() => appStore.theme === 'dark')
const env = computed(() => appStore.env)
// const logInIcon = shallowRef(PersonAddOutline)
// const logOutIcon = shallowRef(PersonRemoveOutline)

async function queryMenu (){
	const res: any = await fetchQueryMenuAPI({ menuPlatform: 1})

	if(!res.success) return
	menuList.value = res.data

	nextTick(() => {
		calcExceededTotalWidth()
	})
}

interface MenuItem {
	menuName: string;
	menuPath: string;
	menuIcon: string;
	menuTipText: string;
	menuIframeUrl: string;
	isJump: boolean;
	isNeedAuth: boolean
}

const menuList = ref<MenuItem[]>([])
const isNeedScroll = ref(false)

onMounted(() => {
	queryMenu()
	useGlobalStore.updateIframeUrl('')
})

const signInStatus = computed(() => Number(authStore.globalConfig?.signInStatus) === 1)


function checkMode() {
  const mode = darkMode.value ? 'light' : 'dark';

  // setThemeAPI({configKey: 'theme', configVal: mode }).then(res => {
		  appStore.setTheme(mode)
 // })
}

function setting() {
  if (!isLogin.value)
    authStore.setLoginDialog(true)

  else
    show.value = true
}
const { isMobile } = useBasicLayout()

const activeRoutePath = computed(() => route.path)

// function toPath(name: string) {
//   router.push({ name })
// }

const mobileSafeArea = computed(() => {
  if (isMobile.value) {
    return {
      paddingBottom: 'env(safe-area-inset-bottom)',
    }
  }
  return {
		height: 'calc(100% - 60px)'
	}
})

const getMobileLayoutClass = computed(() => {
  if (isMobile.value)
    return ['flex-rol', 'w-full', 'border-0']
  return ['flex-col', 'w-sider',]
})

const getIconMobileLayoutClass = computed(() => {
  if (isMobile.value)
    return ['flex', 'flex-rol', 'items-center', 'pt-0', 'w-full']
  return ['flex', 'flex-col', 'pt-1', 'items-center']
})

const iframeSrc = computed(() => useGlobalStore.iframeUrl)

watch(iframeSrc, (newValue) => {
	console.log(newValue);
}, {deep: true,immediate: true})

function handleClickMenu(menu: MenuItem){
	const { menuPath, isJump, menuIframeUrl, isNeedAuth } = menu

	if(isNeedAuth && !isLogin.value){
		message.warning('请先登录后访问！')
		authStore.setLoginDialog(true)
		return;
	}
	useGlobalStore.updateIframeUrl('')
	if(menuPath){
		return router.push({path: menuPath})
	}else{
		if(isJump){
			window.open(menuIframeUrl)
		}else{
			useGlobalStore.updateIframeUrl(menuIframeUrl)
			router.push({ path: '/extend' })
		}
	}
}

function handleSignIn(){
	if(!isLogin.value){
    authStore.setLoginDialog(true)
		return;
	}
	useGlobalStore.updateSignInDialog(true)
}

function calcExceededTotalWidth() {
  if (!track.value) return;
  const { clientHeight = 0, scrollHeight = 0 } = track.value;
  isNeedScroll.value = scrollHeight > clientHeight;
}

watch(
  isMobile,
  (val) => {
    appStore.setSiderCollapsed(val)
  },
  {
    immediate: true,
    flush: 'post',
  },
)
</script>

<template>

  <div class="flex min-w-[80px] pb-2 dark:bg-[#0c0c16] bg-[#fff] dark:border-[#0c0c16]" :class="getMobileLayoutClass" :style="mobileSafeArea">
    <main :class="['flex-1 flex-grow-1 overflow-auto', getIconMobileLayoutClass]" :style="{height: 'calc(100% - 60px)' }" ref="track">
			<NScrollbar>
      <div class="flex h-full flex-col items-center space-y-4">
        <div v-for="item in menuList"
						 :key="item.menuName" class="inline-block"
             :class="isMobile ? 'mt-0' : 'mt-3'"
						 @click="handleClickMenu(item)">
          <NTooltip trigger="hover" placement="right">
            <template #trigger>

							<div class="h-16 w-12 cursor-pointer rounded-xl box-border "
                   :class="[((item.menuIframeUrl && iframeSrc === item.menuIframeUrl) || (item.menuPath && activeRoutePath.includes(item.menuPath))) ?
                   'shadow-[#3076fd] border-[transparent] dark:border-[#1E1E30] bg-gradient-to-l from-[#2F73EC] to-[#5AA8F7] hover:bg-gradient-to-l' : ' border-[#ccc] border-[1px] dark:border-[transparent] bg-white dark:bg-[#1E1E30]']">
                <div class="flex h-full justify-center">

                  <div class="m-auto text-center hover:scale-125 duration-300 transition-all ">
                    <SvgIcon :icon="item.menuIcon" class="text-2xl text-center"
                             :class="[((item.menuIframeUrl && iframeSrc === item.menuIframeUrl) || (item.menuPath && activeRoutePath.includes(item.menuPath))) ||
                             iframeSrc === item.menuIframeUrl ? 'text-[#fff] dark:text-[#fff] ' : '']" />
                 		<p class="mt-1.5 text-[0.7rem] leading-3"
											 :class="[((item.menuIframeUrl && iframeSrc === item.menuIframeUrl) || (item.menuPath && activeRoutePath.includes(item.menuPath))) ?
											  'text-[#fff] dark:text-[#fff] ' : '']">{{item.menuName}}</p>
                  </div>

                </div>
              </div>

            </template>
            <span>{{ item.menuTipText }}</span>
          </NTooltip>
        </div>
      </div>
			</NScrollbar>
    </main>

    <!-- <HoverButton tooltip="全局设置" :placement="isMobile ? 'bottom' : 'right'" :class="isMobile ? 'pb-0' : 'pb-1' " @click="setting">
      <NIcon size="20" color="#555">
        <SvgIcon class="text-2xl" icon="fluent:dark-theme-24-regular" />
      </NIcon>
    </HoverButton> -->
    <!-- <HoverButton v-if="isLogin" tooltip="个人中心" :placement="isMobile ? 'bottom' : 'right'" :class="isMobile ? 'pb-0' : 'pb-8' " @click="toPath('UserCenter')"> -->
    <!-- <SvgIcon icon="icon-park-twotone:personal-collection" /> -->
    <!-- h-[140px] -->

    <div class="flex flex-col justify-between items-center">
      <NTooltip v-if="!isMobile && signInStatus" trigger="hover" placement="right">
        <template #trigger>
          <SvgIcon class="text-xl cursor-pointer mb-5" icon="iconamoon:gift-light" @click="handleSignIn" />
        </template>
        签到奖励
      </NTooltip>

      <NTooltip v-if="!isMobile" trigger="hover" placement="right">
        <template #trigger>
          <SvgIcon class="text-xl cursor-pointer mb-5" :icon="darkMode ? 'ic:outline-dark-mode' : 'entypo:light-up' " @click="checkMode" />
        </template>
        主题切换
      </NTooltip>

<!--      <NTooltip v-if="isLogin" trigger="hover" placement="right">-->
<!--        <template #trigger>-->
<!--          <NAvatar :size="42" :src="avatar" round bordered :fallback-src="defaultAvatar" class="cursor-pointer" @click="toPath('UserCenter')" />-->
<!--        </template>-->
<!--        个人中心-->
<!--      </NTooltip>-->

<!--      <HoverButton v-if="!isLogin" tooltip="登录账户" :placement="isMobile ? 'bottom' : 'right'" :class="isMobile ? 'mb-0' : 'mb-5' " @click="toggleLogin">-->
<!--        <NIcon size="20" color="#555">-->
<!--          <component :is="logInIcon" />-->
<!--        </NIcon>-->
<!--      </HoverButton>-->

    </div>
  </div>
  <Setting v-if="show" v-model:visible="show" />
</template>


<style lang="less" scoped>
.sidebar{
	overflow: hidden;
	width: calc(100% - 5px);
}
.sidebar:hover{
	width: 100%;
	overflow: overlay;
}

.overlay{
	overflow: hidden;
}
.overlay:hover{
	width: 100%;
	overflow: overlay;
}

</style>
