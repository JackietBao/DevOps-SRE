<script setup lang="ts">

import emptyChat from '@/views/chat/components/emptyChat/index.vue'
import {useBasicLayout} from '@/hooks/useBasicLayout'
import {useAuthStore} from '@/store'

const emit = defineEmits<Emit>()
const authStore = useAuthStore()

interface Item {
	[key: number]: any
}

interface Emit {
	(ev: 'run-app', item: Item): void
}

function handleRun(data: any) {
	emit('run-app', data)
}

const {isMobile} = useBasicLayout()
</script>

<template>
	<div :class="[isMobile ? 'mt-2' : '']" class="w-full px-4 py-4 ">
		<div class="w-full md:min-w-[450px]">
			<emptyChat @run="handleRun"/>
		</div>
	</div>
</template>
