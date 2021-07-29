<template>
  <header>
    <img v-if="showLogo !== undefined" :src="logoURL" />
    <action
      v-if="showMenu !== undefined"
      :label="$t('buttons.toggleSidebar')"
      class="menu-button"
      icon="menu"
      @action="openSidebar()"
    />

    <slot />

    <div id="dropdown" :class="{ active: this.$store.state.show === 'more' }">
      <slot name="actions" />
    </div>

    <action
      v-if="this.$slots.actions"
      id="more"
      :label="$t('buttons.more')"
      icon="more_vert"
      @action="$store.commit('showHover', 'more')"
    />

    <div
      v-show="this.$store.state.show == 'more'"
      class="overlay"
      @click="$store.commit('closeHovers')"
    />
  </header>
</template>

<script>
import { logoURL } from "@/utils/constants";

import Action from "@/components/header/Action";

export default {
  name: "header-bar",
  props: ["showLogo", "showMenu"],
  components: {
    Action,
  },
  data: function () {
    return {
      logoURL,
    };
  },
  methods: {
    openSidebar() {
      this.$store.commit("showHover", "sidebar");
    },
  },
};
</script>

<style></style>
