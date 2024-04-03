<script setup lang='ts'>
import type { CSSProperties } from 'vue';
import {computed, ref, watch, nextTick, onMounted} from 'vue';
import type { NumberAnimationInst } from 'naive-ui';
import { NButton, NInput, NLayoutSider, NNumberAnimation, useDialog, useMessage } from 'naive-ui';
import { useRouter } from 'vue-router';
import List from './List.vue';
import {Reply, SvgIcon} from '@/components/common';

import { useAppStore, useAuthStore, useGptsStore, useGlobalStoreWithOut } from '@/store'
import { useBasicLayout } from '@/hooks/useBasicLayout'
const useGlobalStore = useGlobalStoreWithOut();

const router = useRouter();
const appStore = useAppStore();
const gptsStore = useGptsStore();
const authStore = useAuthStore();
const ms = useMessage();
const model3AnimationInstRef = ref<NumberAnimationInst | null>(null);
const model4AnimationInstRef = ref<NumberAnimationInst | null>(null);
const userBalance = computed(() => authStore.userBalance);
const dialog = useDialog();
const { isMobile } = useBasicLayout();
/* 当前选用的模型的扣费类型 1: 普通 2： 高级  */
const activeModelKeyDeductType = computed(() => gptsStore?.activeModelKeyDeductType);
const activeModelKeyPrice = computed(() => gptsStore?.activeModelKeyPrice);

const getMobileClass = computed<CSSProperties>(() => {
	if (isMobile.value) {
		return {
			position: 'fixed',
			zIndex: 50,
			top: '56px',
			height: 'calc(100% - 56px)'
		}
	}
	return {
		background: theme.value ==='dark' ?  '#131323': '',
		height: 'calc(100% - 60px)'
	}
})

const mobileSafeArea = computed(() => {
	if (isMobile.value) {
		return {
			paddingBottom: 'env(safe-area-inset-bottom)',
		}
	}
	return {}
})

const oldUse3Token = ref(0)
const newUse3Token = ref(0)
const oldUse4Token = ref(0)
const newUse4Token = ref(0)

const isSearch = ref(false)
const searchRef = ref(null)

watch(() => authStore.userBalance.useModel3Token, (newVal, oldVal) => {
  oldUse3Token.value = oldVal || 0
  newUse3Token.value = newVal || 0
  model3AnimationInstRef.value?.play()
}, {
  immediate: true,
  flush: 'post',
})

watch(() => authStore.userBalance.useModel4Token, (newVal, oldVal) => {
  oldUse4Token.value = oldVal || 0
  newUse4Token.value = newVal || 0
  model4AnimationInstRef.value?.play()
}, {
  immediate: true,
  flush: 'post',
})

watch(isMobile, (val) => {
		gptsStore.setSiderCollapsed(val)
	},
	{
		immediate: true,
		flush: 'post',
	},
)

const addLoading = ref(false)
const removeLoading = ref(false)

const collapsed = computed(() => gptsStore.siderCollapsed)
const groupKeyWord = ref('')

function handleInputGroupSearch(val: string) {
  groupKeyWord.value = val
  gptsStore.setGroupKeyWord(val)
}

function handleBlurInput(){
	isSearch.value = false;
}

// 提取公共的功能
function checkAuth(){
	if(!authStore.token) {
		ms.info("请先登录")
		addLoading.value = false
		authStore.setLoginDialog(true);
		return false
	}
	 return true
}


/* 新增一个对话 */
async function handleAdd() {

  try {

    if(!checkAuth()) {
      return;
    }

		gptsStore.updateGptsStore(true)
    if (isMobile.value)
			gptsStore.setSiderCollapsed(true)
  }
  catch (error) {
    addLoading.value = false
  }
}

onMounted(() => {
    gptsStore.queryMyGroup()
})

/* 删除全部非置顶聊天 */
async function handleDelGroup() {
	if(!checkAuth()) {
		return;
	}

  dialog.create({
    title: '清空分组',
    content: '是否删除所有非置顶的对话组？',
    positiveText: '确认删除',
    negativeText: '再想想',
    onPositiveClick: async () => {
      await gptsStore.delAllGroup()
    },
  })

}

function handleUpdateCollapsed() {
	gptsStore.setSiderCollapsed(!collapsed.value)
}

const theme = computed(() => appStore.theme)

function handleOpenSearch(){

	if(!checkAuth()) {
		return;
	}

  isSearch.value = true
	nextTick(() => {
		searchRef.value?.focus()
	})
}

</script>

<template>
  <div>
    <NLayoutSider
      :collapsed="collapsed"
      :collapsed-width="0"
      :width="260"
      :show-trigger="false"
      collapse-mode="transform"
      bordered
      :style="getMobileClass"
      @update-collapsed="handleUpdateCollapsed"
    >

      <div class="flex flex-col h-full" :style="mobileSafeArea">
        <main class="flex flex-col h-full flex-1 min-h-0">

          <div class="flex h-14 items-center space-x-2 bg-white px-4 dark:bg-[#131323]">
						<NButton type="primary" :loading="addLoading" @click="handleAdd">
              <SvgIcon icon="ic:round-wifi-find" class="text-xl" />
            </NButton>
            <div class="flex-1">
              <NInput v-model="groupKeyWord"
											ref="searchRef"
											type="text"
											placeholder="对话历史查找"
											@blur="handleBlurInput"
											clearable @input="handleInputGroupSearch" />
            </div>
          </div>

					<div class="flex w-full h-14 items-center space-x-2 bg-white px-4 dark:bg-[#131323]"  v-if="!isSearch">
            <div class=" w-full flex justify-between items-center">

							<NButton  :loading="removeLoading" @click="handleDelGroup">
								清除列表
							</NButton>
              <NButton type="primary" :loading="addLoading" @click="handleAdd">
                更多GPTs
                <SvgIcon icon="ic:round-wifi-find" class="text-xl" />
              </NButton>
            </div>

          </div>

					<!-- 对话分组 显示-->
          <div class="flex-1 min-h-0 pb-4 overflow-hidden">
            <List />
          </div>

<!--          <div class="p-4 w-full border-t dark:border-t-neutral-800 flex flex-col" v-if="!isMobile">-->

<!--            <div class="flex justify-between  w-full">-->
<!--              <Reply></Reply>-->
<!--            </div>-->

<!--          </div>-->

        </main>
      <!-- <Footer /> -->
      </div>
    </NLayoutSider>
    <template v-if="isMobile">
      <div v-show="!collapsed" class="fixed inset-0 z-40 bg-black/40" @click="handleUpdateCollapsed" />
    </template>
  </div>
</template>
