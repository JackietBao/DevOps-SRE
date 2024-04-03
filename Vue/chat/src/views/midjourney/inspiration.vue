<script setup lang="ts">
import {ref} from 'vue'
import Market from '@/views/midjourney/components/market.vue';
import {InspirationTab} from "@/views/midjourney/components/mj-menu";
import {fetchQuerySpireClassify} from "@/api/mjDraw";
import {useBasicLayout} from "@/hooks/useBasicLayout";
const {isMobile} = useBasicLayout();
const marketRef = ref()
const activeTab = ref('')
const tabListRef = ref();
const inspireTypeList = ref<{name: string, value: string}[]>([])
const handleChangeTab = (param: {value: string}) => {
  activeTab.value = param.value;
  marketRef.value.handleQuery(param.value)
}

const handleQueryInspireType = async () => {
  const expand = await fetchQuerySpireClassify({page: 1, size: 100}).then(res => res.data.rows);
  inspireTypeList.value = [{name: "全部", value: ''}].concat(expand)  || InspirationTab;
  marketRef.value.handleQuery(activeTab.value)
}

handleQueryInspireType();


</script>

<template>
   <div class="h-full overflow-hidden" :class="[isMobile ? 'p-0' : 'p-3']">
		  <div class="flex justify-between"  v-if="!isMobile">
				<h2 class="text-lg mb-1.5">AI绘画示例广场</h2>
			</div>
			<!-- tab -->
		 <div class="flex w-full flex-j mb-2 items-center overflow-x-hidden " v-if="!isMobile">

       <div class="relative z-40 transition-300 ease">
         <ul class="flex justify-start items-center" ref="tabListRef">
           <li v-for="(item) in inspireTypeList"
               @click="handleChangeTab(item)"
               class="mr-1 ml-1 w-auto select-none cursor-pointer
						   hover:dark:text-[#fff] hover:text-[#fff]
					     hover:bg-gradient-to-l from-[#2F73EC] to-[#5AA8F7]
						   duration-300 flex rounded-3xl"
               :class="[item.value === activeTab ? 'text-[#fff] bg-gradient-to-l from-[#2F73EC] to-[#5AA8F7]  dark:text-[#fff]': 'bg-[transparent] dark:bg-[transparent]']">
             <div class="pl-3 pr-3 pt-1 pb-1 whitespace-nowrap">{{item.name}}</div>
           </li>
         </ul>
       </div>

     </div>

		 <!-- 规则 -->
		 <div class="flex-1 h-[82vh]">
			 <Market ref="marketRef" :type="activeTab"></Market>
		 </div>

	 </div>
</template>

<style scoped lang="less">

</style>
