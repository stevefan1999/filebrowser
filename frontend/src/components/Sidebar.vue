<template>
  <nav :class="{ active }">
    <template v-if="isLogged">
      <button
        class="action"
        @click="toRoot"
        :aria-label="$t('sidebar.myFiles')"
        :title="$t('sidebar.myFiles')"
        to="/files/"
      >
        <i class="material-icons">folder</i>
        <span>{{ $t("sidebar.myFiles") }}</span>
      </button>

      <div v-if="user.perm.create">
        <button
          :aria-label="$t('sidebar.newFolder')"
          :title="$t('sidebar.newFolder')"
          class="action"
          @click="$store.commit('showHover', 'newDir')"
        >
          <i class="material-icons">create_new_folder</i>
          <span>{{ $t("sidebar.newFolder") }}</span>
        </button>

        <button
          :aria-label="$t('sidebar.newFile')"
          :title="$t('sidebar.newFile')"
          class="action"
          @click="$store.commit('showHover', 'newFile')"
        >
          <i class="material-icons">note_add</i>
          <span>{{ $t("sidebar.newFile") }}</span>
        </button>
      </div>

      <div>
        <button
          class="action"
          @click="toSettings"
          :aria-label="$t('sidebar.settings')"
          :title="$t('sidebar.settings')"
          to="/settings"
        >
          <i class="material-icons">settings_applications</i>
          <span>{{ $t("sidebar.settings") }}</span>
        </button>

        <button
          v-if="authMethod == 'json'"
          id="logout"
          :aria-label="$t('sidebar.logout')"
          :title="$t('sidebar.logout')"
          class="action"
          @click="logout"
        >
          <i class="material-icons">exit_to_app</i>
          <span>{{ $t("sidebar.logout") }}</span>
        </button>
      </div>
    </template>
    <template v-else>
      <router-link
        :aria-label="$t('sidebar.login')"
        :title="$t('sidebar.login')"
        class="action"
        to="/login"
      >
        <i class="material-icons">exit_to_app</i>
        <span>{{ $t("sidebar.login") }}</span>
      </router-link>

      <router-link
        v-if="signup"
        :aria-label="$t('sidebar.signup')"
        :title="$t('sidebar.signup')"
        class="action"
        to="/login"
      >
        <i class="material-icons">person_add</i>
        <span>{{ $t("sidebar.signup") }}</span>
      </router-link>
    </template>

    <p class="credits">
      <span>
        <span v-if="disableExternal">File Browser</span>
        <a
          v-else
          href="https://github.com/filebrowser/filebrowser"
          rel="noopener noreferrer"
          target="_blank"
        >File Browser</a
        >
        <span> {{ version }}</span>
      </span>
      <span
      ><a @click="help">{{ $t("sidebar.help") }}</a></span
      >
    </p>
  </nav>
</template>

<script>
import { mapGetters, mapState } from "vuex";
import * as auth from "@/utils/auth";
import { authMethod, disableExternal, noAuth, signup, version, } from "@/utils/constants";

export default {
  name: "sidebar",
  computed: {
    ...mapState(["user"]),
    ...mapGetters(["isLogged"]),
    active() {
      return this.$store.state.show === "sidebar";
    },
    signup: () => signup,
    version: () => version,
    disableExternal: () => disableExternal,
    noAuth: () => noAuth,
    authMethod: () => authMethod,
  },
  methods: {
    toRoot() {
      this.$router.push({ path: "/files/" }, () => {});
      this.$store.commit("closeHovers");
    },
    toSettings() {
      this.$router.push({ path: "/settings" }, () => {});
      this.$store.commit("closeHovers");
    },
    help() {
      this.$store.commit("showHover", "help");
    },
    logout: auth.logout,
  },
};
</script>
