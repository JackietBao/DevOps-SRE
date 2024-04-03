<script setup lang="ts">
import {computed, ref, reactive} from 'vue'
import {NAvatar} from 'naive-ui'
import {useBasicLayout} from '@/hooks/useBasicLayout'
import {useAuthStore, useGptsStore} from '@/store'

interface Emit {
	(ev: 'run', app: string): void
}
const emit = defineEmits<Emit>()
const authStore = useAuthStore()
const gptsStore = useGptsStore()
const {isMobile} = useBasicLayout()

const groupInfo = computed(() => gptsStore.getChatByGroupInfo())
const showDemoList = computed(() => {

	if(!groupInfo.value || !groupInfo.value ) {
		return [];
	}

	if (groupInfo.value) {
		const demoData = groupInfo.value.demoData?.slice(0,4)

		let newArray: string[] = [];
		let newArray2: string[] = [];

		if (!demoData) return []

		demoData.forEach((item: string, index: number) => {
			index % 2 === 0 ? newArray.push(item): newArray2.push(item)
		})

		return [{children: newArray,}, {children: newArray2,}]
	}
})

const handleGo = (text: string) => {
	emit("run", text)
}

</script>

<template>
  <div :class="[isMobile ? 'mt-2' : '']" class="w-full px-4 py-4 ">
    <div class="w-full md:min-w-[450px] flex justify-center items-center ">
      <div class=" text-center text-sm " :class="[isMobile ? 'w-full' : 'w-1/4 mt-4']">
        <NAvatar :size="80" round :src="groupInfo?.logo"></NAvatar>
        <h2 class="mt-5 text-2xl  ">{{ groupInfo?.title }}</h2>
        <p class="text-center mt-4 text-neutral-400">
          {{
            groupInfo?.desc || ""
          }}</p>
      </div>
    </div>
    <div>

      <div class="flex justify-center items-center " v-if="showDemoList.length"  :class="[isMobile ? 'mt-1' : 'mt-32']">
        <div :class="[isMobile ? 'w-full' : 'w-1/2']">
          <h2 v-if="showDemoList.length">示例提问内容</h2>

          <div v-for="item in showDemoList" class="grid grid-cols-2 lg:grid-cols-2 xl:grid-cols-2 2xl:grid-cols-2 gap-4">
            <div v-for="subItem in item.children">
              <div class="border border-[#E6E6E6] dark:border-[#3B3B56] rounded mt-4 cursor-pointer" @click="handleGo(subItem)">
                <p class="text-md px-2.5 py-2 text-[#AEAEAE] text-center select-none" :class="[isMobile ? 'line-clamp-2 ' : 'line-clamp-1']">{{ subItem }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

    </div>
  </div>
</template>
