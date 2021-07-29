<template>
  <div>
    <p v-if="!isDefault">
      <label for="username">{{ $t("settings.username") }}</label>
      <input
        id="username"
        v-model="user.username"
        class="input input--block"
        type="text"
      />
    </p>

    <p v-if="!isDefault">
      <label for="password">{{ $t("settings.password") }}</label>
      <input
        id="password"
        v-model="user.password"
        :placeholder="passwordPlaceholder"
        class="input input--block"
        type="password"
      />
    </p>

    <p>
      <label for="scope">{{ $t("settings.scope") }}</label>
      <input
        id="scope"
        v-model="user.scope"
        class="input input--block"
        type="text"
      />
    </p>

    <p>
      <label for="locale">{{ $t("settings.language") }}</label>
      <languages
        id="locale"
        :locale.sync="user.locale"
        class="input input--block"
      ></languages>
    </p>

    <p v-if="!isDefault">
      <input
        v-model="user.lockPassword"
        :disabled="user.perm.admin"
        type="checkbox"
      />
      {{ $t("settings.lockPassword") }}
    </p>

    <permissions :perm.sync="user.perm" />
    <commands v-if="isExecEnabled" :commands.sync="user.commands" />

    <div v-if="!isDefault">
      <h3>{{ $t("settings.rules") }}</h3>
      <p class="small">{{ $t("settings.rulesHelp") }}</p>
      <rules :rules.sync="user.rules" />
    </div>
  </div>
</template>

<script>
import Languages from "./Languages";
import Rules from "./Rules";
import Permissions from "./Permissions";
import Commands from "./Commands";
import { enableExec } from "@/utils/constants";

export default {
  name: "user",
  components: {
    Permissions,
    Languages,
    Rules,
    Commands,
  },
  props: ["user", "isNew", "isDefault"],
  computed: {
    passwordPlaceholder() {
      return this.isNew ? "" : this.$t("settings.avoidChanges");
    },
    isExecEnabled: () => enableExec,
  },
  watch: {
    "user.perm.admin": function () {
      if (!this.user.perm.admin) return;
      this.user.lockPassword = false;
    },
  },
};
</script>
