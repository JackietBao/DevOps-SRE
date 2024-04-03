<script setup lang="ts">
import { NSpace, NSkeleton, NEmpty} from 'naive-ui';
import {ref, reactive, watch, computed, onMounted, onUnmounted, onBeforeUnmount} from 'vue'
import taskTab from './components/taskTab.vue'
import taskItem from './components/taskItem.vue'
import {fetchMidjourneyDrawList} from "@/api/mjDraw";
import {useBasicLayout} from "@/hooks/useBasicLayout";
import {useAuthStore} from "@/store";
import {useMidjourneyStore} from "@/store/modules/midjourney";
import {useRoute} from "vue-router";
const route = useRoute()
const activeTab = ref(0);
const totalCount = ref(0);
const loading = ref(false);
const midjourneyStore = useMidjourneyStore();
const showRegionDialog = computed(() => midjourneyStore.openRegionDialog);
const {isMobile} = useBasicLayout();
const isMore = computed(() => totalCount.value > pager.size)
const auth = useAuthStore();
const handleChangeTab = (tab: number) => {

	activeTab.value = tab;

	if (tab !== 6) {
		pager.status = tab;
		pager.favorite = 0;
	} else {
		pager.status = 0;
		pager.favorite = 1; // 1 已收藏 0 未收藏 查询收藏时，状态未全部 111
	}
	handleQueryTaskByPage();
}

const list = ref([])
// {label: '全部', value: 0},
// {label: '执行中', value: 2},
// {label: '已完成', value: 34},
// {label: '已失败', value: 4},
// {label: '我的收藏', value: 6},
// status === 1 ? '等待中' : (status === 2 ? '绘制中' : (status === 3 ? '成功' : (status === 4 ? '失败' : (status === 5 ? '超时' : (status === 6 ? '收藏' : '-')))))
const pager = reactive({
   page: 1,
   size: 20,
	 status: 0,
	 favorite: 0
})

watch(() => showRegionDialog.value, (newValue) => {
   if (!newValue) {
      handleQueryTaskByPage()
   } else {
     loading.value = false;
     timer.value && clearTimeout(timer.value);
   }
})

const timer = ref();

const handleQueryTaskByPage = async () => {
   // loading.value = true;

   const {rows, count}  = await fetchMidjourneyDrawList(pager).then(res => res.data);
   list.value = rows;

   totalCount.value = count;

   // loading.value = false;
   timer.value && clearTimeout(timer.value)

   if (showRegionDialog.value || !route.fullPath.includes("/midjourney/paint-task")) {
       return;
   }

  timer.value = setTimeout(async () => {
    await handleQueryTaskByPage();
  }, 6000);

}


onMounted(() => {

  handleQueryTaskByPage()

  const container: any = document.getElementById('footer')
  const observer = new IntersectionObserver((entries, observer) => {
    entries.forEach((entry) => {
      if (entry.isIntersecting) {
        if (isMore.value) {
          pager.size= pager.size + 10
          handleQueryTaskByPage()
        }
      }
    })
  })
  observer.observe(container)
})

onUnmounted(() => {
  timer.value && clearTimeout(timer.value);
})

onBeforeUnmount(() => {
	auth.updateDrawId({id: ''})
})


</script>

<template>
	<div class="w-full" :class="isMobile ? 'h-full p-2' : ' p-3'" >
		<taskTab @change="handleChangeTab" :active="activeTab"></taskTab>
		<div class="h-full overflow-y-auto pb-2 max-h-[calc(80vh)]">

			<div>
        <div v-if="loading && list.length === 0" class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 gap-4">
          <div v-for="item in 10" class="mb-4  border border-gray-200 dark:border-gray-600 p-2 rounded-xl">
            <n-space vertical >

              <div class="flex justify-between">
                <n-skeleton height="20px" width="48%"   class="mt-1" :sharp="false" />
                <n-skeleton height="20px" width="48%"   class="mt-1" :sharp="false" />
              </div>

              <n-skeleton height="20px" width="100%"  class="mt-1" :sharp="false" />
              <n-skeleton height="180px" width="100%"  class="mt-1" :sharp="false" />

              <n-skeleton height="20px" width="60%"  class="mt-1" :sharp="false" />
              <n-skeleton height="20px" width="40%"  class="mt-1" :sharp="false" />
              <n-skeleton height="20px" width="20%"  class="mt-1" :sharp="false" />

              <div class="flex justify-between">
                <n-skeleton height="20px" width="48%"   class="mt-1" :sharp="false" />
                <n-skeleton height="20px" width="48%"   class="mt-1" :sharp="false" />
              </div>
            </n-space>
          </div>
        </div>
        <div v-else-if="list.length > 0">
          <div class="grid grid-cols-1 lg:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 gap-4">
            <div v-for="item in list" class="">
              <taskItem :task="item" class="mb-4" @fresh="handleQueryTaskByPage"></taskItem>
            </div>
          </div>

        </div>
        <div v-else>
          <div class="flex justify-center items-center">
            <n-empty description="未查询到任务"></n-empty>
          </div>

        </div>
        <div id="footer" ref="containerRef" />
			</div>
		</div>
	</div>
</template>
