<script setup lang="ts">
import {computed} from 'vue'
import Loading from '@/components/base/Loading.vue'
import {NButton, NTooltip, NImage, useDialog, useMessage} from 'naive-ui';
import {SvgIcon} from '@/components/common';
import {
  fetchRemoveMjDraw,
  fetchSetImageFavor,
  fetchAddMjDrawTaskAPI,
} from "@/api/mjDraw";
import axios from "axios";
import {copyText} from "@/utils/format";
import {utcToShanghaiTime} from "@/utils/date";
import {useAuthStore} from "@/store";

const {updateSendSquire, updateDrawId, updatePreview} = useAuthStore();
const downloadUrl = `${import.meta.env.VITE_GLOB_API_URL}/midjourney/download`;
const dialog = useDialog();
const ms = useMessage();
import drawFail from '@/assets/images/draw-fail.png'
import { MidjourneyActionEnum, MidjourneyActionMap, MidJourneyDrawEnum} from "@/constants";
import {useMidjourneyStore} from "@/store/modules/midjourney";

const midjourneyStore = useMidjourneyStore();
const authStore = useAuthStore()
const props = defineProps({
  task: {
    type: Object as any,
    default: () => ({
      prompt: "",
      image: "",
      time: "",
      updatedAt: "",
    })
  },
  chat: {
    type: String,
    default: "task"
  }
})


const emit = defineEmits(['fresh'])

const calcTips = computed(() => (item: any) => {
  const {progress, status} = item

  if (status === 1)
    return '正在排队中...'
  else if (status === 3)
    return '成功'
  else if (status === 2 && progress !== 100)
    return '正在绘制中...'
  else if (status === 2 && progress === 100)
    return '正在存储图片中...'
  else if (status === 4)
    return '绘图失败'
  else if (status === 5)
    return '绘图取消'
  else if (status === 6)
    return '窗口等待'
  else
    return '绘图失败'
})


const calcStatusMsg = computed(() => (status: number) => {
  return status === 1 ? '排队中' : (status === 2 ? '绘制中' : (status === 3 ? '成功' : (status === 4 ? '失败' : (status === 5 ? '超时' : '-'))))
})

const calcTimeMsg = computed(() => (time: number) => {

  if (!time) return ""

  const m = parseInt((time / 60) + '');
  const second = parseInt(String(time % 60));

  return m + '分' + second + '秒'
})

const calcActionMsg = computed(() => (action: number) => {
  return MidjourneyActionMap[action]
})

const scaleIcons = computed(() => {
	const buttons = props.task?.buttons || [];
	const arrU = []; // 放大
	const arrV = []; // 变化
	const arrUV = []; // Upscale Vary
	const arrZCM = []; // Zoom Custom Make
	const arr = [];

	buttons.forEach(item => {
		// 放大
		if (["U1", "U2", "U3", "U4"].includes(item.label)) {
			arrU.push(item);
		}
		// 变换
		else if (["V1", "V2", "V3", "V4"].includes(item.label)) {
			arrV.push(item);
		}
		else if (item.label.startsWith('Upscale') || item.label.startsWith('Vary')) {
			arrUV.push(item);
		}
		else if (item.label) {
			arrZCM.push(item);
		}
		else {
			arr.push(item)
		}
	})
	return [arrU, arrV, arrUV, arrZCM, arr].filter(item=>item.length)
})


// 收藏
const handleFavorite = async (item: { id: number; favorite: number; }) => {
  const {id, favorite} = item;
  const result = await fetchSetImageFavor({id, favorite: favorite === 1 ? 0 : 1}).then(res => res.data)
  result && ms.success(`${favorite === 0 ? '收藏' : "取消收藏"} 成功`);
  emit('fresh')
}


/* 提交放大绘制任务 */
  async function handleAddDrawTask(params: any) {

  const result = await fetchAddMjDrawTaskAPI({type: MidJourneyDrawEnum.IMAGE_ACTION, params}).then(res => res.data)
  updateDrawId({id: result});
  result && ms.success('提交任务成功！');
}

// 获取动作
const getAction = (item: { label: string, customId: string, emoji: string }): MidjourneyActionEnum => {

  // 放大
  if (["U1", "U2", "U3", "U4"].includes(item.label)) {
    return MidjourneyActionEnum.U_SCALE;
  }

  // 变换
  if (["V1", "V2", "V3", "V4"].includes(item.label)) {
    return MidjourneyActionEnum.V_EXPAND;
  }

  // 重新生成no
  if (!item.label && item.emoji) {

    if (item.emoji === '🔄') {
      return MidjourneyActionEnum.REGENERATE;
    }

    if (['⬅️', '➡️', '⬆️', '⬇️'].includes(item.emoji)) {
      return MidjourneyActionEnum.STRETCH;
    }

    // 图生文
    if (["1️⃣", "2️⃣", "3️⃣", "4️⃣"].includes(item.emoji)) {
      return MidjourneyActionEnum.IMAGE_TO_TEXT_ACTION;
    }
  }

  if (item.label.includes('Upscale (4x)')) {
    return MidjourneyActionEnum.UPSCALE_4X
  }

  if (item.label.includes('Upscale (2x)')) {
    return MidjourneyActionEnum.UPSCALE_2X
  }

  if (item.label.includes('Vary (Subtle)')) {
    return MidjourneyActionEnum.VARY_SUBTLE
  }
  if (item.label.includes('Vary (Strong)')) {
    return MidjourneyActionEnum.VARY_STRONG
  }

  if (item.label.includes('Vary (Region)')) {
    return MidjourneyActionEnum.VARY_REGION
  }

  if (item.label.includes('Make Square')) {
    return MidjourneyActionEnum.SQUARE
  }

  if (item.label.includes('Zoom Out 2x')) {
    return MidjourneyActionEnum.ZOOM2X
  }

  if (item.label.includes('Zoom Out 1.5x')) {
    return MidjourneyActionEnum.ZOOM1_5X
  }

  if (item.label.includes('Custom Zoom')) {
    return MidjourneyActionEnum.ZOOM_CUSTOM
  }

  if (item.label.includes('Redo Upscale (Subtle)')) {
    return MidjourneyActionEnum.REDO_UPSCALE_SUBTLE
  }

  if (item.label.includes('Redo Upscale (Creative)')) {
    return MidjourneyActionEnum.REDO_UPSCALE_CREATIVE
  }

  return MidjourneyActionEnum.NONE

}

// 按钮操作
const handleClick = async (task: { prompt: string, mode: number, id: number, imageUrl }, subItem: { label: string, customId: string, emoji: string },) => {

	if (!authStore.isLogin) {
		ms.warning("请先登录")
		authStore.setLoginDialog(true);
		return;
	}

	const action = getAction(subItem);

	await handleAddDrawTask({
		action,
		taskId: task.id,
		customId: subItem.customId,
		messageFlags: task.messageFlags || 0
	})

	emit("fresh")

}

const openRemixModal = (task: any) => {
	// 局部重绘 modal
	if (task.action === 9) {
		midjourneyStore.openRegionModal(true, {
			customId: task.customId,
			taskId: task.id,
			mode: task.mode
		});
	}
	// remix modal
	else {
		authStore.updateRemixModal(true, {
			prompt: '',
			taskId: task.id,
			customId: task.customId,
		});
	}
}


/* 复制 */
const handleCopyPrompt = (item: { fullPrompt: string }) => {
  const {fullPrompt} = item
  copyText({text: fullPrompt})
  ms.success('复制prompt完成！')
}

/**
 * 直接从 imageUrl 下载图片
 */
async function handleDownload({ imageUrl }: { imageUrl: string }) {

	if (!imageUrl) {
		ms.warning('图片未生成完成暂不能下载')
		return
	}

	dialog.info({
		title: '下载图片',
		content: '是否确认下载当前图片',
		positiveText: '下载',
		negativeText: '取消',
		onPositiveClick: async () => {
			const filename = imageUrl.split('/');
			const response = await axios.post(downloadUrl, { url: imageUrl }, { responseType: 'blob' })
			const blob = new Blob([response.data], { type: response.headers['content-type'] })
			const urlObject = window.URL.createObjectURL(blob)
			const link = document.createElement('a')
			link.href = urlObject
			link.download = filename[filename.length - 1]
			link.click()
		}
	})
}

// 删除
const handleRemove = (item: { id: number; favorite: number }) => {

  const {favorite} = item;

  dialog.create({
    type: "info",
    title: '确认删除',
    content: `${favorite === 1 ? '当前任务已被收藏，删除绘画任务会导致收藏失效且不可找回，确定删除' : '删除绘画任务不可找回，确定删除'}？`,
    positiveText: '确定',
    negativeText: '取消',
    maskClosable: false,
    onPositiveClick: async () => {
      const result = await fetchRemoveMjDraw({id: item.id}).then(res => res.data)
      result && ms.success("删除成功")
      emit("fresh")
    },
    onNegativeClick: () => {

    }

  })

}


// 发布到广场
const handlePublishSquare = async (param: { id: number }) => {
  updateSendSquire(true, param.id)
}

const imageProp = {
  referrerPolicy: "https://cdn.discordapp.com",
}

// 发布到广场
const handlePreview = async (param: { id: number }) => {
  updatePreview(true, param.id)
}

</script>

<template>
  <div class="border  box-border border-[#eaeaef] dark:border-[#444456] rounded-xl  min-h-fit flex-1">
    <div class="pt-2 px-1 pb-1">

      <div class="flex justify-between items-center">

        <p class="border pt-0.5 px-1 text-[12px] scale-95 rounded">
          {{ calcActionMsg(task.action) }} - {{ calcStatusMsg(task.status) }} - {{
            calcTimeMsg(task.durationSpent) || ""
          }}
        </p>

        <div class="flex justify-between ">
          <NTooltip placement="top-start" trigger="hover" v-if="task.status === 3">
            <template #trigger>
              <SvgIcon icon="bi:send" @click="handlePublishSquare(task)" class="text-[14px] mx-1 cursor-pointer"/>
            </template>
            <div>
              发布到广场
            </div>
          </NTooltip>

          <!--          <NTooltip placement="top-start" trigger="hover" v-if="task.action === 1 && task.status === 3">-->
          <!--            <template #trigger>-->
          <!--              <SvgIcon icon="mingcute:scissors-line" class="text-[14px] mx-1 cursor-pointer"/>-->
          <!--            </template>-->
          <!--            <div>-->
          <!--              拆分图片-->
          <!--            </div>-->
          <!--          </NTooltip>-->

          <NTooltip placement="top-start" trigger="hover">
            <template #trigger>
              <SvgIcon icon="ph:copy-light" @click="handleCopyPrompt(task)" class="text-[14px] mx-1 cursor-pointer"/>
            </template>
            <div>
              复制最终执行提示词
            </div>
          </NTooltip>

          <NTooltip placement="top-start" trigger="hover" v-if="task.status === 3">
            <template #trigger>
              <SvgIcon icon="icons8:down-round" @click="handleDownload(task)" class="text-[14px] mx-1 cursor-pointer"/>
            </template>
            <div>
              下载图片
            </div>
          </NTooltip>


          <NTooltip placement="top-start" trigger="hover">
            <template #trigger>
              <SvgIcon @click="handleRemove(task)" icon="ep:delete" class="text-[14px] mx-1 cursor-pointer"/>
            </template>
            <div>
              删除任务
            </div>
          </NTooltip>

          <NTooltip placement="top-start" trigger="hover">
            <template #trigger>
              <SvgIcon @click="handleFavorite(task)" :icon="task.favorite === 0 ? 'uit:favorite' : 'uis:favorite'"
                       class="text-[14px] mx-1 cursor-pointer"/>
            </template>
            <div>
              收藏此任务
            </div>
          </NTooltip>

        </div>
      </div>

      <div class="my-2">
        <NTooltip placement="top" trigger="hover" :width="320">
          <template #trigger>
            <h3 class="whitespace-nowrap overflow-hidden text-ellipsis" style="height: 22px;">{{ task.prompt }}</h3>
          </template>
          <div>
            {{ task.prompt }}
          </div>
        </NTooltip>
      </div>

      <div class="relative z-0 w-auto flex justify-center align-center min-h-[260px] overflow-hidden contain"
           :class="[chat === 'chat-history' ? 'h-[320px]' : 'h-[260px] ']">

        <div v-if="task.status === 4"
             class=" bg-[#f0f0f0] dark:bg-[#454567] h-full w-full flex flex-col justify-center items-center">
          <img class="w-auto object-contain px-24 " :src="drawFail" alt="">
        </div>

				<div v-else-if="task.status === 5"
             class=" bg-[#f0f0f0] dark:bg-[#454567] h-full w-full flex flex-col justify-center items-center">
          <img class="w-auto object-contain px-24 " :src="drawFail" alt="">
        </div>

				<div v-else-if="task.status === 6"
             class=" bg-[#f0f0f0] dark:bg-[#454567] h-full w-full flex flex-col justify-center items-center">
          <img class="w-auto object-contain px-24 " :src="drawFail" alt="">

        </div>

        <Loading :progress="task.progress" v-else-if="!task.imageUrl && !task.progress"></Loading>

        <div v-else class="flex items-center justify-center flex-col">
          <NImage class="w-full cursor-pointer bg-[#f9f9f9] dark:bg-[transparent] h-full flex justify-center"
                  :preview-src="task.imageUrl"
									 object-fit="contain"
                  :src="task.imageUrl || ''"></NImage>
        </div>

      </div>
    </div>

    <div class="bg-[#f5f5f5] dark:bg-[#3E3E64] px-2 py-2 rounded-br-xl  rounded-bl-xl h-[130px]">

      <div v-if="!task.imageUrl && task.progress === null" class="flex justify-center h-[80px]">

				<div v-if="calcTips(task)" class="progress">
					<span>
						{{ calcTips(task) }}
					</span>
					<template v-if="task.status === 4">
						<span>：</span>
						<span style="color: red;">{{ task.failReason }}</span>
					</template>

					<n-button type="primary" v-if="task.status === 6" style="margin-left: 12px;"
						@click="openRemixModal(task)"
					>
						补充描述词
					</n-button>
        </div>
      </div>

      <div class="flex items-center justify-center h-[80px]"
           v-if="task.progress !== null && task.progress >= 0 && task.progress < 100">
        <p class="text-center flex items-center">
          <SvgIcon icon="eos-icons:bubble-loading" class="mr-1">
          </SvgIcon>
          进度：{{ task.progress + '%' }}
        </p>
      </div>

      <div v-if="task.progress === 100" class="h-[80px]">
        <div v-for="(item) in scaleIcons" class="flex justify-start items-center py-0.5 mb-1">
          <div class="flex justify-start">
            <template v-for="_item in item">
              <n-tooltip trigger="hover">
                <template #trigger>
                  <button :disabled="task.status !== 3" :class="[task.status !== 3 ? 'cursor-not-allowed' : '']" class="text-[14px]  border border-gray-300
                       dark:border-gray-600 hover:text-white
                       hover:bg-gradient-to-l  from-[#2F73EC] to-[#5AA8F7]
                       px-1 flex items-center justify-center rounded"
                          style="margin: 0 4px"
                          @click="handleClick(task, _item)">
                    <p v-if="_item.label" class="text-[12px] pt-[2px] line-clamp-1 text-left">{{ _item.label }}</p>
                    <p v-else>{{ _item.emoji }}</p>
                  </button>
                </template>
                <p>{{ _item.label || _item.emoji }}</p>
              </n-tooltip>
            </template>
          </div>
        </div>
      </div>

      <div class="flex justify-between text-[12px] mt-4 select-none">
        <div><p class="text-[#3144F1] dark:text-[#fff]">时间：{{ utcToShanghaiTime(task.createdAt) }}</p></div>
        <div><p>{{task.mode === 1 ? '快速模式' : task.mode === 2 ? '慢速模式' : task.mode === 3 ? 'Turbo模式' : '模式错误'}}</p></div>
      </div>
    </div>

  </div>
</template>

