<script setup lang="ts">
import {computed, ref,onMounted, watch, onBeforeUnmount} from 'vue'
import textConfigure from './components/configure.vue'
import chatMessage from './components/chatMessage.vue'
import chatConfigure from './components/chatConfigure.vue'
import chatHistory from './components/chatHistory.vue'
import MobileText2Image from './mobile/text2image.vue'
import {useAuthStore} from "@/store";
import {fetchGetDraw} from "@/api/mjDraw";
import {useBasicLayout} from "@/hooks/useBasicLayout";
import {MidjourneyStatusEnum} from "@/views/midjourney/components/mj-menu";
import Suggestion from "@/views/midjourney/components/suggestion.vue";
const {isMobile} = useBasicLayout();
const auth = useAuthStore();

const jobId = computed(() => auth.drawId);

const remixModalOpen = computed(() => auth.openRemixModal)

// 获取当前绘图任务
const handleGetCurrentDraw = async (jobId: string) => {
	if (!jobId) {
		return;
	}

	currentTask.value =  await fetchGetDraw({jobId}).then(res => res.data).then(res => {
    if (res.status === 6) {
      auth.updateRemixModal(true, {
				prompt: '',
				taskId: res.id,
        customId: res.customId,
      });
      drawTimer.value && clearTimeout(drawTimer.value);
    }
    return res;
  })
  drawTimer.value && (currentTask.value.status >= 4) && clearTimeout(drawTimer.value);
}

const handleGetTask = () => {
  handleGetCurrentDraw(jobId.value);
  handlePullByTimes()
}

const currentTask = ref({
  status: 0
})

const drawTimer = ref();

const handlePullByTimes = () => {
  drawTimer.value && clearTimeout(drawTimer.value);
  drawTimer.value = setTimeout(() => {

    if (currentTask.value?.status >= MidjourneyStatusEnum.DRAWED) {
      drawTimer.value && clearTimeout(drawTimer.value);
      return;
    }

    handlePullByTimes();
    handleGetCurrentDraw(jobId.value)
   }, 3000)
}

watch(remixModalOpen, (newValue, oldValue) => {
	if (!newValue && jobId.value) {
		handleGetTask()
	}
})

onBeforeUnmount(() => {
  drawTimer.value && clearTimeout(drawTimer.value);
  drawTimer.value = null;
  auth.updateDrawId({id: ''})
})



</script>

<template>

	<div v-if="isMobile">
		<MobileText2Image></MobileText2Image>
	</div>

	<div class="text-image-container flex flex-row"  v-if="!isMobile">
		<div class="w-full pt-2 px-20 h-[calc(100% - 60px)] flex flex-col justify-between" >

			<!--	中间的对话框	-->
			<div class="h-full flex justify-start items-center"  >

				<div class="w-full" v-if="!jobId ">
					<Suggestion></Suggestion>
				</div>

				<template v-if="jobId">
					<chatHistory :task="currentTask"  @refresh="handleGetTask"></chatHistory>
				</template>
			</div>

			<chatConfigure ></chatConfigure>
			<chatMessage @fresh="handleGetTask" class="m-full"></chatMessage>

		</div>
	   <div>
			 <textConfigure ref="configRef" ></textConfigure>
		 </div>
	</div>
</template>
