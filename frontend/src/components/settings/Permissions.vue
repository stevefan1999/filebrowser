<template>
  <div>
    <h3>{{ $t("settings.permissions") }}</h3>
    <p class="small">{{ $t("settings.permissionsHelp") }}</p>

    <p>
      <input v-model="admin" type="checkbox" />
      {{ $t("settings.administrator") }}
    </p>

    <p>
      <input v-model="perm.create" :disabled="admin" type="checkbox" />
      {{ $t("settings.perm.create") }}
    </p>
    <p>
      <input v-model="perm.delete" :disabled="admin" type="checkbox" />
      {{ $t("settings.perm.delete") }}
    </p>
    <p>
      <input v-model="perm.download" :disabled="admin" type="checkbox" />
      {{ $t("settings.perm.download") }}
    </p>
    <p>
      <input v-model="perm.modify" :disabled="admin" type="checkbox" />
      {{ $t("settings.perm.modify") }}
    </p>
    <p v-if="isExecEnabled">
      <input v-model="perm.execute" :disabled="admin" type="checkbox" />
      {{ $t("settings.perm.execute") }}
    </p>
    <p>
      <input v-model="perm.rename" :disabled="admin" type="checkbox" />
      {{ $t("settings.perm.rename") }}
    </p>
    <p>
      <input v-model="perm.share" :disabled="admin" type="checkbox" />
      {{ $t("settings.perm.share") }}
    </p>
  </div>
</template>

<script>
import { enableExec } from "@/utils/constants";

export default {
  name: "permissions",
  props: ["perm"],
  computed: {
    admin: {
      get() {
        return this.perm.admin;
      },
      set(value) {
        if (value) {
          for (const key in this.perm) {
            this.perm[key] = true;
          }
        }

        this.perm.admin = value;
      },
    },
    isExecEnabled: () => enableExec,
  },
};
</script>
