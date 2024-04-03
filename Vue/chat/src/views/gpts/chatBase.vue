<script setup lang='ts'>
import type {Ref} from 'vue'
import {computed, nextTick, reactive, onMounted, onUnmounted, ref, watch} from 'vue'
import {NButton, NInput, NTooltip, useDialog, useMessage} from 'naive-ui'
import EmptyLog from '@/views/gpts/components/emptyLog/index.vue'
import CopyRight from "@/components/common/Copyright/index.vue";
import RecordDialog from "@/layout/components/RecordDialog.vue";
import HeaderComponent from "@/views/gpts/components/Header/index.vue";
import type {App} from '@/views/appStore/components/helpter'
import html2canvas from 'html2canvas'
import {useRoute} from 'vue-router'
import {Message} from './components'
import {useScroll} from './hooks/useScroll'
import {useCopyCode} from './hooks/useCopyCode'
import {useGptsChat} from './hooks/useGptsChat'
import {useUsingContext} from './hooks/useUsingContext'
import {useUsingNetwork} from './hooks/useUsingNetwork'

import {SvgIcon} from '@/components/common'
import {useBasicLayout} from '@/hooks/useBasicLayout'
import {useAppStore, useAuthStore, useGptsStore, useGlobalStoreWithOut} from '@/store'
import {fetchChatAPIProcess} from '@/api'
import {t} from '@/locales'
import {router} from '@/router'

import {fetchChatTailAPI} from "@/api/config";
import {useUpload} from "@/views/chat/hooks/useUpload";
import {AxiosProgressEvent} from "axios";
import AiGptsCom from "@/views/gpts/components/gpts/aiGptsCom.vue";

const useGlobalStore = useGlobalStoreWithOut()
const authStore = useAuthStore()
const route = useRoute()
let controller = new AbortController()

const dialog = useDialog()
const ms = useMessage()
const appStore = useAppStore()
const bottomContainer = ref()
const gptsStore = useGptsStore()
const isStreamIn = ref(false)
const typingStatusEnd = ref(true) // 打字效果是否完成

const {addGroupChat, updateGroupChat, updateGroupChatSome} = useGptsChat()

const appId = computed(() => route.query.appId as string)
const tradeStatus = computed(() => route.query.trade_status as string)
const token = computed(() => route.query.token as string)

useCopyCode()

const {isMobile} = useBasicLayout()
const {scrollRef, scrollToBottom, scrollToBottomIfAtBottom} = useScroll()
const {usingContext, toggleUsingContext} = useUsingContext()
const {usingNetwork, toggleUsingNetwork} = useUsingNetwork()

const dataSources = computed(() => gptsStore.chatList)
const gptsSources = computed(() => gptsStore.groupList)
const activeModel = computed(() => gptsStore.activeModelName)
const isCanUpload = computed(() => gptsStore.canUpload);
const isCanAudio = computed(() => gptsStore.canAudio);

const {selectFile, upload} = useUpload();

/* 当前所有的ai回复信息列表 方便拿到上下文 */
const conversationList = computed(() => dataSources.value.filter(item => (!item.inversion && !item.error)))

const prompt = ref<string>('')
const loading = ref<boolean>(false)
const inputRef = ref<Ref | null>(null)
const firstScroll = ref<boolean>(true)
const chatTail = ref<string>('')

/* 当前选中的对话组id */
const activeGroupId = computed(() => gptsStore.active)

/* 当前对话组的详细信息 */
const activeGroupInfo = computed(() => gptsStore.groupList.find((item: any) => item.uuid === activeGroupId.value))

/* 当前选用的模型的类型 1： openai  2: 百度  */
const activeModelKeyType = computed(() => Number(gptsStore?.activeModelKeyType))

const buttonDisabled = computed(() => {
  return loading.value || !prompt.value || prompt.value.trim() === '' || !typingStatusEnd.value
})

watch(activeGroupId, (val) => {
      if (val)
        firstScroll.value = true
      if (inputRef.value && !isMobile.value)
        inputRef.value?.focus()
    },
    {immediate: true},
)


watch(dataSources, (val) => {
      if (val.length === 0)
        return
      if (firstScroll.value) {
        firstScroll.value = false
        scrollToBottom()
      }
    },
    {immediate: true},
)

async function handleOpenRecord() {
  useGlobalStore.updateRecordDialog(true)
}

// 上传文件
async function handleUploadFile(mode: number | File = 1) {

  const purpose = "assistants";

  let file: File | number;

  if (mode === 1) {
    file = await selectFile()

    if (!file) {
      ms.error("选择文件错误")
      return;
    }
  } else {
    file = mode;
  }

  useGlobalStore.updateUploadDialog(true)

  const result = await upload({purpose, file: <File>file}, (progressEvt: AxiosProgressEvent) => {
    const {progress, loaded, total} = progressEvt;

    if (progress && Number(progress.toFixed(2)) * 100 <= 99) {
      useGlobalStore.updateUploadProgress(Number(progress.toFixed(2)) * 100);
    }

  })

  if (result.data) {
    useGlobalStore.updateUploadProgress(100);
    useGlobalStore.updateUploadDialog(false)

    prompt.value = mode === 1 ? '请分析一下上传的这个文件 ' + result.data : '请将语音转换成文本 ' + result.data
  }

  if (mode !== 1) {
    useGlobalStore.updateRecordDialog(false)
  }

}

function handlePrompt(val: string) {
  prompt.value = val;
  handleSubmit()
}

function handleUpdateModel(val: number) {
  useGlobalStore.updateModel(val)
}

/* 发送消息 */
async function handleSubmit(index?: number) {

  if (gptsStore.groupList.length === 0 || loading.value || !typingStatusEnd.value) return;
  let message = ''

  /* 如果有index就是重新生成 */
  if (index && typeof index === 'number') {
    const {requestOptions} = dataSources.value[index]
    message = requestOptions?.prompt ?? ''
  }

  onConversation(message)
}

/* 按钮发送消息 */
async function onConversation(msg?: string) {
  let message = msg || prompt.value

  if (loading.value)
    return

  if (!message || message.trim() === '')
    return

  controller = new AbortController()

  /* 虚拟增加一条用户记录 */
  addGroupChat({
    dateTime: new Date().toLocaleString(),
    text: message,
    inversion: true,
    error: false,
    conversationOptions: null,
    requestOptions: {prompt: message, options: null},
  })

  scrollToBottom()

  loading.value = true
  prompt.value = ''

  let options: any = {groupId: +activeGroupId.value, type: 'gpts', usingNetwork: usingNetwork.value}
  const lastContext = conversationList.value[conversationList.value.length - 1]?.conversationOptions

  if (lastContext && usingContext.value && !usingNetwork.value) {
    options = {...lastContext, ...options}
  }


  /* 虚拟增加一条ai记录 */
  addGroupChat({
        dateTime: new Date().toLocaleString(),
        text: 'AI思考中',
        loading: true,
        inversion: false,
        error: false,
        conversationOptions: null,
        requestOptions: {prompt: message, options: {...options}},
      },
  )

  scrollToBottom()

  let timer: any = null;
  let cacheResText = ''; // 部分模型是分批返回文字 需要合并 全部的内容
  let currentText = ''; // 当前使用的部分
  let i: number = 0;
  let data: any = null;
  isStreamIn.value = true;
  let userBanance: any = {};

  useGlobalStore.updateIsChatIn(true)

  /* 匀速输出 */
  try {
    const fetchChatAPIOnce = async () => {

      timer = setInterval(async () => {

        cacheResText[i] && currentText.length < cacheResText.length && (currentText += cacheResText[i])

        if (cacheResText[i]) {

          typingStatusEnd.value = false
          /* 实时更新消息 */
          updateGroupChat(
              dataSources.value.length - 1,
              {
                dateTime: new Date().toLocaleString(),
                text: currentText,
                inversion: false,
                usage: data?.detail?.usage,
                error: false,
                loading: true,
                conversationOptions: {conversationId: data?.conversationId, parentMessageId: data?.id},
                requestOptions: {prompt: message, options: {...options}},
              },
          )
          scrollToBottomIfAtBottom()
          i++
        }

        /* 更新完了结束 */
        if (!isStreamIn.value && currentText.length >= cacheResText.length) {

          i = 0
          currentText = ''
          timer && clearInterval(timer)
          timer = null;

          // 修改loading状态
          typingStatusEnd.value = true

          updateGroupChatSome(dataSources.value.length - 1, {loading: false}) // 更新最后一条的loading状态

          useGlobalStore.updateIsChatIn(false)

          /* 如果拿到了余额信息  更新 */
          if (userBanance) {
            authStore.updateUserBanance(userBanance)
          }

          /* 当为2说明第一次对话 则修改一个对话组的默认title，取回答记录前十个字 如果长度不到十个 拿全部 */
          // if (dataSources.value.length === 2 && !activeGroupInfo?.value?.appId) {
          // 	const title = dataSources.value[1].text.length > 15 ? dataSources.value[1].text.slice(0, 15) : dataSources.value[1].text
          // 	// await gptsStore.updateGroupInfo({groupId: +activeGroupId.value, title})
          // 	await gptsStore.queryMyGroup()
          // }

          scrollToBottomIfAtBottom()
        }
      }, 16)

      // 流式接入数据
      await fetchChatAPIProcess<Chat.ConversationResponse>({
        prompt: message,
        appId: activeGroupInfo.value ? activeGroupInfo.value.appId : 0,
        options,
        signal: controller.signal,
        onDownloadProgress: ({event}) => {

          const xhr = event.target;
          const {responseText} = xhr

          /* 这种解析只对openai有效 其他的会漏掉前面的字 */
          const lastIndex = responseText.lastIndexOf('\n', responseText.length - 2)

          let chunk = responseText
          if (lastIndex !== -1) {
            chunk = responseText.substring(lastIndex)
          }

          try {
            data = JSON.parse(chunk)

          } catch (error) {
            // TODO 如果出现类似超时错误 会连接上次的内容一起发出来导致无法解析  后端需要处理 下
            console.log('parse data erro from openai: ', error, chunk);
            if (chunk.includes('OpenAI timed out waiting for response')) {
              console.log('绘画超时了~')
              ms.warning('会话超时了、告知管理员吧~~~')
            }
          }

          // 显示消耗的token
          try {
            cacheResText = data.text;
            isStreamIn.value = !data.is_end;

            if (data?.detail?.usage) {
              typingStatusEnd.value = data.is_end;
              userBanance = data?.userBanance;
            }


          } catch (error) {
            console.log(error);
          }
        }
      })
    }

    await fetchChatAPIOnce()
  } catch (error: any) {
    console.log('chat error catch', error);

    useGlobalStore.updateIsChatIn(false)
    clearInterval(timer)
    isStreamIn.value = false

    if (error.code === 402 || error?.message.includes('余额不足') || error?.message.includes('免费额度已经使用完毕'))
      useGlobalStore.updateVipDialog(true)

    let errorMessage = error?.message ?? t('common.wrong')
    if (errorMessage === 'Request failed with status code 401')
      errorMessage = '非法操作、请先登录后再进行问答使用！'

    if (error?.message.includes('canceled')) {
      updateGroupChatSome(dataSources.value.length - 1, {loading: false})
      scrollToBottomIfAtBottom()
      setTimeout(() => {
        authStore.getUserBalance()
      }, 200)
      return
    }

    const currentChat = dataSources.value[dataSources.value.length - 1]

    if (currentChat?.text && currentChat.text !== '') {
      updateGroupChatSome(
          dataSources.value.length - 1,
          {
            text: `${currentChat.text === 'AI思考中' ? '' : currentChat.text}\n[${errorMessage}]`,
            error: false,
            loading: false,
          },
      )
      return
    }

    updateGroupChat(
        dataSources.value.length - 1,
        {
          dateTime: new Date().toLocaleString(),
          text: errorMessage,
          inversion: false,
          error: true,
          loading: false,
          conversationOptions: null,
          requestOptions: {prompt: message, options: {...options}},
        },
    )
    scrollToBottomIfAtBottom()
  } finally {
    console.log('finally', loading.value);

    loading.value = false
    isStreamIn.value = false
  }
}

/* 重新生成 */
async function onRegenerate(index: number) {

  if (loading.value) {
    return
  }

  controller = new AbortController()

  const {requestOptions} = dataSources.value[index]

  let message = requestOptions?.prompt ?? ''

  let options: any = {groupId: +activeGroupId.value, usingNetwork: usingNetwork.value}

  if (requestOptions.options) {
    options = {...requestOptions.options, groupId: +activeGroupId.value}
  }

  loading.value = true

  /* 虚拟增加一条用户记录 */
  addGroupChat(
      {
        dateTime: new Date().toLocaleString(),
        text: message,
        inversion: true,
        error: false,
        conversationOptions: null,
        requestOptions: {prompt: message, ...options},
      },
  )

  /* 虚拟增加一条ai记录 */
  addGroupChat(
      {
        dateTime: new Date().toLocaleString(),
        text: '',
        loading: true,
        inversion: false,
        error: false,
        conversationOptions: null,
        requestOptions: {prompt: message, options: {...options}},
      },
  )

  scrollToBottom()

  try {

    let lastText = ''
    const fetchChatAPIOnce = async () => {

      await fetchChatAPIProcess<Chat.ConversationResponse>({
        prompt: message,
        appId: activeGroupInfo.value ? activeGroupInfo.value.appId : 0,
        options,
        signal: controller.signal,

        onDownloadProgress: ({event}) => {
          const xhr = event.target
          const {responseText} = xhr
          // Always process the final line
          const lastIndex = responseText.lastIndexOf('\n', responseText.length - 2)
          let chunk = responseText

          if (lastIndex !== -1)
            chunk = responseText.substring(lastIndex)
          try {

            const data = JSON.parse(chunk)
            /* 拿到剩余次数额度 */
            const userBanance = data?.detail?.userBanance
            if (userBanance)
              authStore.updateUserBanance(userBanance)

            /* 实时更新消息 */
            updateGroupChat(
                dataSources.value.length - 1,
                {
                  dateTime: new Date().toLocaleString(),
                  text: lastText + (data.text ?? ''),
                  inversion: false,
                  usage: data?.detail?.usage,
                  error: false,
                  loading: true,
                  conversationOptions: {conversationId: data.conversationId, parentMessageId: data.id},
                  requestOptions: {prompt: message, options: {...options}},
                },
            )
          } catch (error) {
            console.log(error);
            //
          }
        },
      })
      updateGroupChatSome(dataSources.value.length - 1, {loading: false})
    }
    await fetchChatAPIOnce()
  } catch (error: any) {

    if (error?.message.includes('canceled')) {
      updateGroupChatSome(dataSources.value.length - 1, {loading: false})
      setTimeout(() => authStore.getUserBalance(), 200)
      return
    }

    if (error.code === 402 || error?.message.includes('余额不足'))
      useGlobalStore.updateVipDialog(true)

    let errorMessage = error?.message ?? t('common.wrong')
    if (errorMessage === 'Request failed with status code 401')
      errorMessage = '非法操作、请先登录后再进行问答使用！'

    updateGroupChat(
        dataSources.value.length - 1,
        {
          dateTime: new Date().toLocaleString(),
          text: errorMessage,
          inversion: false,
          error: true,
          loading: false,
          conversationOptions: null,
          requestOptions: {prompt: message, ...options},
        },
    )

  } finally {
    loading.value = false
    isStreamIn.value = false
  }
}

/* 有appId的特殊处理 */
async function handleAppId(appId: number) {
  const id = +appId
  // await router.replace({name: 'Chat', query: {}})
  const isHasGroup = gptsStore.groupList.find((item: any) => item.appId === id)
  if (!isHasGroup) {
    await gptsStore.addNewChatGroup(id)
    await gptsStore.queryMyGroup()
  } else {
    await gptsStore.setActiveGroup(isHasGroup.uuid)
  }

  if (prompt.value) {
    handlePrompt(prompt.value)
  }

}

/* 处理三方对接跳转 */
async function otherLoginByToken(token: string) {
  authStore.setToken(token)
  await router.replace({name: 'Chat', query: {}})
  ms.success('账户登录成功、开始体验吧！')
  await authStore.getUserInfo()
}

/* 支付回调处理 */
function handleRefresh() {
  if (tradeStatus.value.toLowerCase().includes('success')) {
    ms.success('感谢你的购买、祝您使用愉快~', {duration: 5000})
    authStore.getUserInfo()
    router.replace({name: 'Chat', query: {}})
  } else {
    ms.error('您还没有购买成功哦~')
  }
}

/* 导出 */
function handleExport() {
  if (loading.value) {
    return
  }


  const d = dialog.create({
    title: t('chat.exportImage'),
    content: t('chat.exportImageConfirm'),
    positiveText: t('common.yes'),
    negativeText: t('common.no'),
    onPositiveClick: async () => {
      try {
        d.loading = true
        const ele = document.getElementById('image-wrapper')
        const canvas = await html2canvas(ele as HTMLDivElement, {
          useCORS: true,
        })
        const imgUrl = canvas.toDataURL('image/png')
        const tempLink = document.createElement('a')
        tempLink.style.display = 'none'
        tempLink.href = imgUrl
        tempLink.setAttribute('download', 'chat-shot.png')
        if (typeof tempLink.download === 'undefined')
          tempLink.setAttribute('target', '_blank')

        document.body.appendChild(tempLink)
        tempLink.click()
        document.body.removeChild(tempLink)
        window.URL.revokeObjectURL(imgUrl)
        d.loading = false
        ms.success(t('chat.exportSuccess'))
        return Promise.resolve(true)
      } catch (error: any) {
        ms.error(t('chat.exportFailed'))
      } finally {
        d.loading = false
      }
    },
  })
}

/* 删除 */
function handleDelete({chatId}: Chat.Chat) {

  if (loading.value) {
    ms.warning("对话中，不可进行删除操作");
    return
  }

  dialog.warning({
    title: t('chat.deleteMessage'),
    content: t('chat.deleteMessageConfirm'),
    positiveText: t('common.yes'),
    negativeText: t('common.no'),
    onPositiveClick: () => {
      gptsStore.deleteChatById(chatId)
    },
  })
}

function handleClear() {

  if (loading.value)
    return

  dialog.create({
    title: t('chat.clearChat'),
    content: t('chat.clearChatConfirm'),
    positiveText: t('common.yes'),
    negativeText: t('common.no'),
    onPositiveClick: async () => {
      await gptsStore.clearChatByGroupId()
      ms.success('删除当前页面对话完成')
    },
  })
}

function handleEnter(event: KeyboardEvent) {
  if (!isMobile.value) {
    if (event.key === 'Enter' && !event.shiftKey) {
      event.preventDefault()
      handleSubmit()
    }
  } else {
    if (event.key === 'Enter' && event.ctrlKey) {
      event.preventDefault()
      handleSubmit()
    }
  }
}

function handleStop() {
  if (loading.value) {
    controller.abort()
    loading.value = false
    isStreamIn.value = false
    typingStatusEnd.value = true
  }
}

const placeholder = computed(() => {
  if (isMobile.value)
    return t('chat.placeholderMobile')
  return t('chat.placeholder')
})

// 获取小尾巴功能
const getChatTail = () => {
  fetchChatTailAPI(['dialogueTail']).then(res => {
    chatTail.value = res.data;
  })
}

const handleChat = async (data: App) => {
  if (data.id && data.demoData) {
    prompt.value = data.demoData;
    await handleAppId(data.id);
    return;
  }
  prompt.value = ""
}

onMounted(async () => {

  if (appId.value) {
    await handleAppId(Number(appId.value))
  }
  if (token.value) {
    await otherLoginByToken(token.value)
  }

  if (tradeStatus.value) {
    handleRefresh()
  }

  await nextTick(async () => {
    await gptsStore.queryActiveChatLogList()
    await scrollToBottom()
    if (inputRef.value && !isMobile.value)
      inputRef.value?.focus()
  })
  getChatTail()
})
const darkMode = computed(() => appStore.theme === 'dark')
const recordDialog = computed(() => useGlobalStore.recordDialog)

const getMobileStyle = computed(() => {
	if (!isMobile.value)
		return {height: 'calc(100% - 60px)'}
})

const btnSize = computed(() => {
	if (!isMobile.value)
		return 'small'
	return 'tiny'
})

const gpts = reactive({
  query: ""
})

const gptsItemRef = ref()

// 搜索
const search = (type?: string) => {
  if (type === 'clear') {
    gpts.query = ""
  }

  gptsItemRef.value.search(gpts.query);
}

onUnmounted(() => {
  if (loading.value)
    controller.abort()
})

</script>

<template>
  <div class=" flex flex-col bg-white dark:bg-[#1f1f38]" :style="getMobileStyle">

    <!-- 上传文件弹窗 -->
    <RecordDialog :visible="recordDialog" @submit="handleUploadFile"/>

    <!--头部显示标题 和导出按钮 -->
    		<HeaderComponent
    			:using-context="usingContext"
    			:dark-mode="darkMode"
    			@export="handleExport"
    			@toggle-using-context="toggleUsingContext"
    			@clear="handleClear"
    		/>

    <main class="flex-1 overflow-hidden">

      <div id="scrollRef" ref="scrollRef" class="relative h-full overflow-hidden overflow-y-auto scroll-smooth">
        <div
            id="image-wrapper"
            class="w-full max-w-screen-2xl m-auto dark:bg-[#1F1F38] h-full"
            :class="[isMobile ? 'p-2' : 'p-4']"
        >

          <div :class="[isMobile ? 'h-full' : 'h-[84vh]']" v-if="!gptsSources.length">
            <div class="flex justify-between items-center  w-full">
              <div class="pr-4">GPTS</div>
              <div class="max-w-[500px]">
                <n-input round placeholder="GPTs名字(enter搜索)" class="w-full" @clear="search('clear')" clearable
                         v-model:value="gpts.query" @keyup.enter="search()">
                  <template #prefix>
                    <SvgIcon icon="uil:search"/>
                  </template>
                  <template #suffix>
                    <div class="cursor-pointer" @click="search()">搜索</div>
                  </template>
                </n-input>
              </div>

            </div>

            <div class="w-full h-[1px] bg-[#E4E4E4] dark:bg-[#444456] mt-3"></div>

            <AiGptsCom ref="gptsItemRef" :q="gpts.query"/>

          </div>

          <template v-if="!dataSources.length && gptsSources.length > 0">
            <EmptyLog @run="handlePrompt"></EmptyLog>
          </template>

          <template v-else>

            <!--新增小尾巴	-->
            <div>
              <Message
                  :chatTail="chatTail"
                  v-for="(item, index) of dataSources"
                  :key="item.chatId"
                  :date-time="item.dateTime"
                  :text="item.text"
                  :inversion="item.inversion"
                  :error="item.error"
                  :loading="item.loading"
                  @regenerate="handleSubmit(index)"
                  @delete="handleDelete(item)"
              />
              <div class="sticky bottom-0 left-0 flex justify-center">
                <NButton v-if="loading" @click="handleStop">
                  <template #icon>
                    <SvgIcon icon="ri:stop-circle-line"/>
                  </template>
                  停止输出
                </NButton>
              </div>
              <div ref="bottomContainer" class="bottom"/>
            </div>
          </template>
        </div>
      </div>
    </main>


    <footer class="footer-dialog" v-if="gptsSources.length > 0">
      <div :class="['m-auto max-w-screen-2xl', isMobile ? 'px-0 pt-1' : 'px-4 py-2']">

        <div class="flex justify-between">
          <div class="flex mb-2">

            <NTooltip trigger="hover" :disabled="isMobile" v-if="isCanUpload">
              <template #trigger>
                <NButton round :size="btnSize" @click="handleUploadFile(1)"  style="diplay: block; margin-right: 1rem;">
							    <span class="flex justify-center items-center">
										 <SvgIcon class="text-lg" style="width: 1em;height: 1em" icon="octicon:link-16"/>
                     <span>上传</span>
								  </span>
                </NButton>
              </template>
              上传文件
            </NTooltip>

            <NTooltip trigger="hover" :disabled="isMobile" v-if="isCanAudio">
              <template #trigger>
                <NButton :size="btnSize"  @click="handleOpenRecord" round class="ml-4">
                     <span class="flex justify-center items-center">
								      	 <SvgIcon class="text-lg" style="width: 1em;height: 1em" icon="streamline:voice-mail"/>
                       <span>语音</span>
								       </span>
                </NButton>
              </template>
              语音助手
            </NTooltip>
          </div>

          <div class="mb-2" v-if="dataSources.length > 0">
            <NTooltip trigger="hover" :disabled="isMobile">
              <template #trigger>
                <NButton round :size="btnSize"  @click="handleClear" >
							    <span class="flex justify-center items-center ">
										 <SvgIcon class="text-lg" style="width: 0.8em;height: 0.8em; margin-right: 4px;" icon="uiw:delete"/>
                     <span>清空</span>
								  </span>
                </NButton>
              </template>
              清空聊天记录
            </NTooltip>
          </div>
        </div>

        <div class="flex items-stretch space-x-2">
          <div class="relative flex-1">

            <NInput
                ref="inputRef"
                v-model:value="prompt"
                type="textarea"
                class="pb-10"
                clearable
                :placeholder="placeholder"
                :autosize="{ minRows: isMobile ? 1 : 2, maxRows: isMobile ? 3 : 4 }"
                @keypress="handleEnter"
            >
            </NInput>

            <div class="absolute bottom-2   right-2 w-[28px]" >
              <div class="flex justify-end items-center">
                <NButton type="primary" size="small"
                         :disabled="buttonDisabled" round @click="handleSubmit">
                  <template #icon>
														<span class="dark:text-white">
															<SvgIcon icon="icon-park-outline:send"/>
														</span>
                  </template>
									发送
                </NButton>
              </div>
            </div>
          </div>
        </div>
      </div>
    </footer>

    <CopyRight v-if="!isMobile" class="footer-copyright-bar"></CopyRight>


  </div>
</template>

<style lang="less">
.message-container {
  flex: 3;
}
</style>
