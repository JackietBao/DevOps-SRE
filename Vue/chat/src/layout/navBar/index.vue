<script setup lang="ts">
import Logo from '../siderBar/Logo.vue';
import MjTabMenu from './MjTabMenu.vue';
import defaultAvatar from "@/assets/avatar.png";
import {NAvatar,NTooltip, NPopover, NButton} from "naive-ui";
import { SvgIcon, SIcon } from "@/components/common";
import {computed, ref} from "vue";
import {useAppStore, useAuthStore, useGlobalStore, useGlobalStoreWithOut} from "@/store";
import {useBasicLayout} from "@/hooks/useBasicLayout";
import {useRoute, useRouter} from "vue-router";
import {fetchSignLogAPI} from "@/api/signin";
import {utcToShanghaiTime} from "@/utils/date";

const globalStore = useGlobalStore();
const globalWithout = useGlobalStoreWithOut()
const appStore = useAppStore()
const authStore = useAuthStore();
const isLogin = computed(() => authStore.isLogin);
const avatar = computed(() => authStore.userInfo.avatar);
const username = computed(() => authStore.userInfo.username);
const userBalance = computed(() => authStore.userBalance);

const { isMobile } = useBasicLayout()
const router = useRouter()
function toggleLogin() {
	if (isLogin.value)
		authStore.logOut()

	else
		authStore.setLoginDialog(true)
}

function toPath(name: string) {
	router.push({ name })
}

function handleSignIn(){
	if(!isLogin.value){
		authStore.setLoginDialog(true)
		return false
	}
	globalWithout.updateSignInDialog(true)
}

function logOut() {
	authStore.logOut()
	router.push('/')
}

function handleGo(type: string) {
	globalStore.updateUserInfoDialog(true, type)
}


const signInData = ref([])
function signed(month: number, date: number) {
	const str = `${new Date().getFullYear()}-${month.toString().padStart(2, '0')}-${date.toString().padStart(2, '0')}`
	const res: any = signInData.value.find((item: any) => item.signInDate === str)
	return !res ? false : res?.isSigned
}

const hasSignedInToday = computed(() => {
	const month = new Date().getMonth() + 1
	const date = new Date().getDate()
	return signed(month, date)
})

async function getSigninLog() {
	signInData.value = await fetchSignLogAPI().then(res => res.data || [])
}

const theme =  computed(() => appStore.getTheme());

const fillColor = computed(() => {
	return {'fill': theme.value === 'dark' ? '#fff' : '#555'}
})

const handleGoVip = () => {
	// router.push("/pay")
  handleGoMember();
}

// 会员中心
const handleGoMember = async () => {
  if(authStore.checkAuth()) {
    globalStore.updateVipDialog(true)
  }
}


const route = useRoute();

const showTab = computed(() => {
  return route.fullPath.includes("midjourney")
})

</script>

<template>
    <div class="navbar-container relative z-50 h-[60px] bg-[#FcFcFc] dark:bg-[#131323] dark:border-[#3a3a40] border-#efeff5-800 flex justify-between items-center">

			<div class="cw-logo ml-4 h-[60px] flex items-center">
					 <Logo  />
           <MjTabMenu v-if="showTab"></MjTabMenu>
			 </div>

			<div class="mr-4 flex items-center flex-end">
				<NTooltip trigger="hover" placement="bottom">
					<template #trigger>
						<div class="flex justify-between items-center mr-3 cursor-pointer select-none" @click="globalStore.updateNoticeDialog(true)">
							<SvgIcon class="text-xl mr-2 text-gray-500 dark:text-gray-300" icon="fluent:speaker-2-20-regular"></SvgIcon>
						</div>
					</template>
					<p>系统公告</p>
				</NTooltip>

				<NTooltip trigger="hover" placement="bottom">
					<template #trigger>
						<div class="flex justify-between items-center mr-3 cursor-pointer select-none" @click="globalStore.updateFeedbackDialog(true)">
							<SvgIcon class="text-xl mr-2  text-gray-500 dark:text-gray-300" icon="icon-park-outline:email-push"></SvgIcon>
						</div>
					</template>
					<p>反馈我们</p>
				</NTooltip>

				<NTooltip trigger="hover" placement="bottom">
					<template #trigger>
						<div class="flex justify-between items-center mr-3 cursor-pointer select-none" @click="handleGoMember">
							<SvgIcon class="text-xl mr-2 text-gray-500 dark:text-gray-300" icon="mingcute:vip-2-line"></SvgIcon>
						</div>
					</template>
					<p>会员中心</p>
				</NTooltip>

					<div v-if="isLogin" class=" px-2.5 cursor-pointer rounded-md" >
						<n-popover trigger="click"
											 raw
											 arrow-class="dark:bg-[#2D2D4F] bg-[#fafafa"
											 class="dark:bg-[#2D2D4F] bg-[#fafafa] rounded-xl px-2 py-2"
											 :on-update:show="getSigninLog"
						>
							<template #trigger>
								<div class="flex justify-between items-center px-2 py-1 text-white" >
                  <NAvatar :size="40"
                           :src="avatar"
                           round
                           bordered
                           :fallback-src="defaultAvatar"
                           class="cursor-pointer"
                  />
                  <div class="flex justify-center items-center text-gray-500 dark:text-gray-300">
                    <p class="select-none ml-1">{{username}}</p>
                    <SvgIcon class="text-xl" icon="mingcute:down-small-fill" />
                  </div>

								</div>
							</template>

							<div class="w-[300px] px-2 py-2 rounded-2xl p-4 ">
								 <div class="flex justify-between items-center w-full">
									 <div class="flex items-center flex-start">
 											<NAvatar :size="46"
															 :src="avatar"
															 round
															 bordered
															 :fallback-src="defaultAvatar"
															 class="cursor-pointer"
															 />
										 <p class="text-xl ml-2">{{username}}</p>
									 </div>

									 <n-button  size="small" @click="handleSignIn">
										 {{hasSignedInToday ? '已签到' : '签到'}}
									 </n-button>
								 </div>

								<div class="rounded-2xl my-6 flex justify-between items-center py-2 px-4 dark:bg-gradient-to-l from-[#4D8ADD] to-[#BCD8FF]  bg-gradient-to-l">
										<div >
											<h2 class="text-xl dark:text-[#114386]">开通Vip</h2>
											<p v-if="userBalance?.expirationTime" class="text-xs dark:text-[#114386]">到期时间：{{utcToShanghaiTime(userBalance?.expirationTime)}}</p>
											<p v-else class="text-xs dark:text-[#114386]">未开通，点击了解特权</p>
										</div>
										<div class="cursor-pointer" @click="handleGoVip">
											<p class="dark:text-[#114386] flex flex-start">
												{{ userBalance.expirationTime ? '已开通' : '立即开通' }}
												<SvgIcon class="text-xl" icon="mdi:menu-right" />
											</p>
										</div>
								</div>

								<div class="flex justify-between mt-6">
										<div class="flex flex-col justify-center cursor-pointer" @click="handleGo('baseInfo')">
											<div class="flex justify-center">
												<SIcon name="baseInfo" :style="fillColor"></SIcon>
											</div>
											<p class="text-[14px]">基本信息</p>
										</div>
										<div class="flex flex-col justify-center cursor-pointer" @click="handleGo('invite')">
											<div class="flex justify-center">
												<SIcon name="invite" :style="fillColor"></SIcon>
											</div>
											<p class="text-[14px]">邀请好友</p>
										</div>
										<div class="flex flex-col justify-center cursor-pointer"  @click="handleGo('exchange')">
											<div class="flex justify-center">
												<SIcon name="ticket" :style="fillColor"></SIcon>
											</div>
											<p class="text-[14px]">权益兑换</p>
										</div>
								</div>

								<div class="mt-6">
									<div class="flex flex-col w-1/5 cursor-pointer"  @click="logOut">
										<div class="flex justify-center">
											<SIcon name="logout" :style="fillColor"></SIcon>
										</div>
										<p>退出登录</p>
									</div>

								</div>

							</div>
						</n-popover>

					</div>

					<div class="bg-[#3187F1] px-2.5 cursor-pointer rounded-md" v-if="!isLogin"  @click="toggleLogin">
            <div class="flex justify-between items-center px-2 py-1 text-[#f0f0f0]" >
             <SvgIcon icon="iconoir:user" class="text-lg"></SvgIcon>
              <p class="select-none">未登录</p>
            </div>

					</div>

				</div>
		</div>
</template>

<style scoped lang="less">
  .navbar-container {
		box-shadow: 0px 4px 10px rgba(0,0,0,0.1);
	}
</style>
