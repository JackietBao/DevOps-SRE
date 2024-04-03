<script setup lang='ts'>
import { computed, watch } from 'vue'
import { NLayout, NLayoutContent } from 'naive-ui'
import LeftSider from './components/sider/index.vue'
import ChatBase from './chatBase.vue'
import AiGpts from './components/gpts/aiGpts.vue'
import { useBasicLayout } from '@/hooks/useBasicLayout'
import {useAppStore, useAuthStore, useChatStore, useGptsStore} from '@/store'

const appStore = useAppStore()
const gptsStore = useGptsStore()
const authStore = useAuthStore()
const { isMobile } = useBasicLayout()

const isLogin = computed(() => authStore.isLogin)

watch(isLogin, async (newVal, oldVal) => {
  if (newVal && !oldVal) {
    await gptsStore.queryMyGroup()
  }

})

const getMobileClass = computed(() => {
  if (isMobile.value)
    return ['rounded-none', 'shadow-none']
  return ['rounded-md', 'shadow-md', 'dark:border-neutral-800']
})

const getContainerClass = computed(() => {
  return [
    'h-full',
    // { 'pl-[260px]': !isMobile.value && !collapsed.value },
  ]
})


</script>

<template>
  <div class="h-full dark:bg-[#131323] transition-all">
    <div class="h-full overflow-hidden" :class="getMobileClass">

		  <!-- gpts store-->
			<AiGpts></AiGpts>

      <NLayout class="z-40 transition" :class="getContainerClass" :has-sider="true" sider-placement="left">

        <!-- 左侧侧侧对话页面 的侧边栏 -->
        <LeftSider class="h-full" />

        <!-- 聊天-->
        <ChatBase class="w-full"/>

      </NLayout>
    </div>
  </div>
</template>
