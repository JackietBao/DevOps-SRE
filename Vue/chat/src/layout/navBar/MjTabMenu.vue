<script setup lang="ts">
import {NTabPane, NTabs,} from 'naive-ui';
import {computed, ref} from "vue";
import {MjMenu} from "@/views/midjourney/components/mj-menu";
import {useRoute, useRouter} from "vue-router";
const router = useRouter();
const route = useRoute();

const menuList = ref(MjMenu)
const activeRoutePath = computed(() =>  route.path as string);

const handleChangeTab = (menuPath: string) => {
  router.push(`${menuPath}`);
}

</script>

<template>
  <n-tabs
      v-model:value="activeRoutePath"
      class="card-tabs flex justify-center items-center"
      default-value="signin"
      size="medium"
      animated
      @update-value="handleChangeTab"
      pane-wrapper-style="margin: 0 -4px"
      pane-style="padding: 0 4px 0 4px; box-sizing: border-box;"
  >
    <n-tab-pane :name="item.menuPath" :tab="item.menuName" v-for="item in menuList" :key="item.key"></n-tab-pane>
  </n-tabs>
</template>
