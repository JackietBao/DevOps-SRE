<script setup lang='ts'>
import { NButton, NIcon, useMessage, NModal, NResult, NTabPane, NTabs} from 'naive-ui'
import {computed, ref,  } from 'vue'
import {CloseOutline} from '@vicons/ionicons5'
import Phone from './Login/Phone.vue'
import Email from './Login/Email.vue'
import Wechat from './Login/Wechat.vue'
import {useAuthStore, useAppStore} from '@/store'
import {SvgIcon} from '@/components/common'
import {useBasicLayout} from "@/hooks/useBasicLayout";
defineProps<Props>()
let timer: any
const authStore = useAuthStore()
const appStore = useAppStore()
const activeCount = ref(false)
const wxLoginUrl = ref('')
const sceneStr = ref('')
const {isMobile} = useBasicLayout()
const theme = computed(() => appStore.theme)

const ms = useMessage();

const registerSendStatus = computed(() => {
	return Number(authStore.globalConfig.registerSendStatus)
})

const registerSendModel3Count = computed(() => {
	return Number(authStore.globalConfig.registerSendModel3Count)
})

const registerSendModel4Count = computed(() => {
	return Number(authStore.globalConfig.registerSendModel4Count)
})

const registerSendDrawMjCount = computed(() => {
	return Number(authStore.globalConfig.registerSendDrawMjCount)
})

const wechatRegisterStatus = computed(() => Number(authStore.globalConfig.wechatRegisterStatus) === 1)
const phoneRegisterStatus = computed(() => Number(authStore.globalConfig.phoneRegisterStatus) === 1)
const emailRegisterStatus = computed(() => Number(authStore.globalConfig.emailRegisterStatus) === 1)

const disabledReg = computed(() => {
	return !wechatRegisterStatus.value && !phoneRegisterStatus.value && !emailRegisterStatus.value
})

const activeTab = ref(wechatRegisterStatus.value ? 'wechat' : 'email')

// watchEffect(() => {
//   if (wechatRegisterStatus.value) {
//     activeTab.value = ""
//   }
//   if (phoneRegisterStatus.value) {
//     activeTab.value = ""
//   }
// })

interface Props {
	visible: boolean
}

const registerTips = computed(() => (`首次认证：赠送${registerSendModel3Count.value}积分基础模型 | ${registerSendModel4Count.value}积分高级模型 | ${registerSendDrawMjCount.value}积分绘画`))


function handleCloseDialog() {
	clearInterval(timer)
	wxLoginUrl.value = ''
	sceneStr.value = ''
	activeCount.value = false
}

const closeColor = computed(() => {
	return theme.value === 'light' ? '#555' : '#fff'
})

const handleGoLogin = (mode: string) => {
	activeTab.value = mode;
}

const closeLogin = () => {
  if (!authStore.isLogin) {
    ms.warning("请先登录系统后再使用");
    return;
  }
  authStore.setLoginDialog(false);
}

</script>

<template>
	<NModal :show="visible"
					style="max-width: 550px; min-height: 600px;"
					class="dark:bg-[#33334D] bg-[#F2F2F2] "
					:class="[isMobile ? 'px-2 w-full' : 'px-20']"
					:on-after-leave="handleCloseDialog">
		<div class=" rounded " :class="[isMobile ? 'h-[100vh]' : 'p-5']">

			<div class="absolute top-3 right-3 cursor-pointer z-30" v-if="!isMobile" @click="closeLogin">
				<NIcon size="20" :color="closeColor">
					<CloseOutline/>
				</NIcon>
			</div>


			<div v-if="disabledReg">
				<NResult
					size="small"
					status="403"
					title="网站已经关闭注册通道"
					description="请联系管理员开通吧"
				>
					<template #footer>
						<NButton @click="authStore.setLoginDialog(false)">
							知道了
						</NButton>
					</template>
				</NResult>
			</div>

			<div v-else>

				<!-- register -->
				<div class="mx-auto"  v-if="!wechatRegisterStatus || activeTab !== 'wechat'" :class="[isMobile ? 'mt-4 w-auto' : 'w-auto']">
					<NTabs  justify-content="center" animated v-model:value="activeTab">
						<NTabPane v-if="emailRegisterStatus" name="email" tab="账号登录"></NTabPane>
						<NTabPane v-if="phoneRegisterStatus" name="phone" tab="手机号登录"></NTabPane>
					</NTabs>
				</div>

				<div v-if="wechatRegisterStatus && activeTab === 'wechat'">
					<h3 class="text-xl text-center mb-5">微信扫码登录</h3>
					<Wechat/>
				</div>

				<div :class="[isMobile ? 'mt-4' : '']" >

<!--
					<div v-if="isMobile" class="flex flex-col  items-center">
						<img :src="logoPath ? logoPath : logo"   class="w-[60px] cursor-pointer px-0 dark:border-[#ffffff17] border-#ebebeb-400" alt="">
&lt;!&ndash;						<NAvatar :size="60"&ndash;&gt;
&lt;!&ndash;										 :src="logoPath"&ndash;&gt;
&lt;!&ndash;										 round&ndash;&gt;
&lt;!&ndash;										 bordered&ndash;&gt;
&lt;!&ndash;										 :fallback-src="logo"&ndash;&gt;
&lt;!&ndash;										 class="cursor-pointer"&ndash;&gt;
&lt;!&ndash;						/>&ndash;&gt;
						<p class="text-xl w-auto mt-1">{{sideName || "GoMaxAi"}}</p>

					</div>
-->

					<Email v-if="emailRegisterStatus && activeTab === 'email'" />
					<Phone v-if="phoneRegisterStatus && activeTab === 'phone'"/>
				</div>

				<div v-if="!disabledReg" class="w-full flex items-center justify-center mt-10">
					<div class="flex items-center justify-around w-2/3">

						<div v-if="wechatRegisterStatus && activeTab !== 'wechat'"
								 class="select-none cursor-pointer text-[#fff] bg-[#4D8ADD] dark:bg-[#4D4D7A] px-4 py-2 flex justify-center items-center rounded"
								 @click="handleGoLogin('wechat')">
							<SvgIcon icon="ic:twotone-wechat"></SvgIcon>
							<p class="ml-1">微信登录</p>
						</div>

						<div v-if="emailRegisterStatus && activeTab !== 'email'"
								 class="select-none cursor-pointer text-[#fff] bg-[#4D8ADD] dark:bg-[#4D4D7A] px-4 py-2 flex justify-center items-center rounded"
								 @click="handleGoLogin('email')">
							<SvgIcon icon="ic:baseline-email"></SvgIcon>
							<p class="ml-1">账号登录</p>
						</div>

						<div v-if="phoneRegisterStatus && activeTab !== 'phone'"
								 class="select-none cursor-pointer text-[#fff] bg-[#4D8ADD] dark:bg-[#4D4D7A] px-4 py-2 flex justify-center items-center rounded"
								 @click="handleGoLogin('phone')">
							<SvgIcon icon="fluent:phone-16-filled"></SvgIcon>
							<p class="ml-1">手机登录</p>
						</div>

					</div>
				</div>

				<div class="mt-6">
					<p class="text-center text-[12px]">注册登录即表示阅读并同意

						<router-link to="/agreement" target="_blank">
							 <span>
								 《用户协议》
							 </span>
						</router-link>

						<router-link to="/policy" target="_blank">
							 <span>
								《隐私政策》
							 </span>
						</router-link>
					</p>
				</div>

				<div v-if="registerSendStatus && !disabledReg"
						 class="mt-5 border rounded px-2 py-2 text-[12px] dark:border-[#737373] border-[##A7A7A7] ">
					<p class="text-center select-none text-[12px]">{{ registerTips }}</p>
				</div>

			</div>


		</div>
	</NModal>
</template>
