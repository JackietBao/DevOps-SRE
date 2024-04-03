<script setup lang="ts">

import {  useChatStore, useGptsStore } from '@/store';
import {ref, computed, reactive, watch, onMounted} from 'vue';
import {useMessage, NButton, NImage, NTag, NScrollbar} from 'naive-ui';
import { SvgIcon } from '@/components/common';
import { useRouter } from 'vue-router';
import {fetchQueryGptsByPageAPI, fetchQueryGptsGroupAPI} from "@/api/gpts";

const router = useRouter()
const ms = useMessage();
const chatStore = useChatStore()
const gptStore = useGptsStore()
const emit = defineEmits(['close','toq']);
const pp = defineProps<{q:string}>();
const gptsList =  ref<{logo: string,  desc: string, modelName: string, useCount: number}[]>([])
const st = ref({loadPage:false, q:'', tab:'',search:false, hasMore: true});

const tag= ref(<{id: number, groupName: string}[]>[]);
const activeTag= ref(0);


//  const gptUrl = homeStore.myData.session.gptUrl??'https://gpts.ddaiai.com/open/gpts';
const go = async(item: any)=>{
	// await gptStore.upgradeModelConfig(item) // 切换模型，创建对话
	await gptStore.addNewChat(item.id)
  gptStore.updateGptsStore(false)
	ms.success('切换成功！');
}

const pageLoad = async () =>{
		st.value.tab = '';
    st.value.loadPage= true;
    pager.page += 1;
    await	handleQueryList();
    // const gptUrl= `https://gpts.ddaiai.com/open/gptsapi/list/${ gptsPageList.value.length}`;
    st.value.loadPage= false;

}

// 搜索
const search = async (q:string ) => {

    st.value.tab = 'search';
    st.value.search= true;
	  pager.groupId = 0;
	  pager.modelName = q;
  	await	handleQueryList();
    // const gptUrl= `https://gpts.ddaiai.com/open/gptsapi/search?q=${ st.value.q }`;
    st.value.search= false;
}

const pager = reactive({
	 size: 20,
 	 page: 1,
	 groupId: 0,
	 modelName: "",
});

const goSearch = async (q: { id: number })=>{
	activeTag.value = q.id
	pager.groupId = q.id;
	st.value.hasMore = true;
	await	handleQueryList();
}

onMounted(() => {
	st.value.hasMore = true;
	handleQueryGptGroup();
	handleQueryList()
})

const handleQueryList = async () => {
	const list  = await fetchQueryGptsByPageAPI(pager).then(res => res.data.rows || []);

	if (pager.page > 1 && list.length === 0) {
		pager.page -= 1;
		st.value.hasMore = false;
		ms.info("未查询到更多Gpt");
		return;
	}

  if (gptsList.value.length === 0 || activeTag.value !== 0 || st.value.tab === 'search') {
    gptsList.value = list;
  } else {
    gptsList.value = list.concat(gptsList.value);
		st.value.hasMore = true;
  }
}

const handleQueryGptGroup = async () => {
  tag.value = await fetchQueryGptsGroupAPI({page: 1, size: 20}).then(res => [{id: 0, groupName: "全部"}, ...res.data.rows] || []);
}

defineExpose({
	search
})

</script>
<template>

<div class="w-full h-full p-4 overflow-y-hidden">
	<div class="w-full dark:bg-transparent" >
		<NScrollbar x-scrollable>
			<div class="flex items-center justify-start line-clamp-1 pb-4"  >
				<div class="m-1 cursor-pointer select-none" v-for="v in tag" @click="goSearch(v)">
					<div class=" py-1 px-4 rounded-full  border border-gray-200 dark:border-gray-700" :class="[v.id === activeTag ?
							'bg-gradient-to-r from-[#3076ED] to-[#54A2F5] hover:bg-gradient-to-r text-white' : 'text-[#5A5A5A] dark:text-[#fff]  bg-[#FBFBFB] dark:bg-[#1F1F38]']">
						<p>{{ v.groupName}}</p>
					</div>
				</div>
			</div>
		</NScrollbar>
	</div>

  <div class="h-[92%] overflow-y-auto">
    <div class="grid grid-cols-1 gap-5 md:grid-cols-2 lg:grid-cols-3"  >
      <div @click="go(v)" v-for="v in gptsList" class="group relative flex gap-3 rounded-2xl bg-[#f9f9f9] dark:bg-[#313155] p-5  cursor-pointer ">
        <div class="min-w-0 flex-1">
          <div class="flex justify-between items-center ">
            <h3 class=" transition  text-lg font-semibold line-clamp-1 "> {{ v.modelName }}</h3>
            <n-tag type="success" size="small" round v-if="v.useCount && (+v.useCount)> 0">
              <div class="flex items-center text-[#FA6406]"><SvgIcon icon="mdi:hot"  ></SvgIcon>{{ v.useCount }}</div>
            </n-tag>
          </div>
          <div class="mt-0.5 text-zinc-400 text-md line-clamp-2">{{ v.desc }}</div>

        </div>
        <NImage :src="v.logo" :preview-disabled="true" lazy class="group-hover:scale-[130%] duration-300 shrink-0 overflow-hidden bg-base object-cover rounded-full bc-avatar w-[80px] h-[80px]">
          <template #placeholder>
            <div class="w-full h-full justify-center items-center flex"  >
              <SvgIcon icon="line-md:downloading-loop" class="text-[60px] text-green-300"></SvgIcon>
            </div>
          </template>
        </NImage>
      </div>
    </div>

    <div class="h-full flex items-center justify-center"  v-if="gptsList.length === 0">
      未查询到更多GPT
    </div>

		<div class="flex items-center justify-center py-5">
			<div  v-if="st.loadPage">
				<div class="w-full h-full justify-center items-center flex"  >
					<SvgIcon icon="line-md:downloading-loop" class="text-[60px] text-green-300"></SvgIcon>
				</div>
			</div>
			<NButton @click="pageLoad()" v-if="gptsList.length && !st.loadPage && st.hasMore">加载更多</NButton>
		</div>

	</div>



</div>
</template>
