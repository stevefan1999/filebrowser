<template>
  <div class="row">
    <div class="column">
      <form class="card" @submit="updateSettings">
        <div class="card-title">
          <h2>{{ $t("settings.profileSettings") }}</h2>
        </div>

        <div class="card-content">
          <p>
            <input v-model="hideDotfiles" type="checkbox" />
            {{ $t("settings.hideDotfiles") }}
          </p>
          <p>
            <input v-model="singleClick" type="checkbox" />
            {{ $t("settings.singleClick") }}
          </p>
          <h3>{{ $t("settings.language") }}</h3>
          <languages
            :locale.sync="locale"
            class="input input--block"
          ></languages>
        </div>

        <div class="card-action">
          <input
            :value="$t('buttons.update')"
            class="button button--flat"
            type="submit"
          />
        </div>
      </form>
    </div>

    <div class="column">
      <form v-if="!user.lockPassword" class="card" @submit="updatePassword">
        <div class="card-title">
          <h2>{{ $t("settings.changePassword") }}</h2>
        </div>

        <div class="card-content">
          <input
            v-model="password"
            :class="passwordClass"
            :placeholder="$t('settings.newPassword')"
            name="password"
            type="password"
          />
          <input
            v-model="passwordConf"
            :class="passwordClass"
            :placeholder="$t('settings.newPasswordConfirm')"
            name="password"
            type="password"
          />
        </div>

        <div class="card-action">
          <input
            :value="$t('buttons.update')"
            class="button button--flat"
            type="submit"
          />
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import { mapMutations, mapState } from "vuex";
import { users as api } from "@/api";
import Languages from "@/components/settings/Languages";

export default {
  name: "settings",
  components: {
    Languages,
  },
  data: function () {
    return {
      password: "",
      passwordConf: "",
      hideDotfiles: false,
      singleClick: false,
      locale: "",
    };
  },
  computed: {
    ...mapState(["user"]),
    passwordClass() {
      const baseClass = "input input--block";

      if (this.password === "" && this.passwordConf === "") {
        return baseClass;
      }

      if (this.password === this.passwordConf) {
        return `${baseClass} input--green`;
      }

      return `${baseClass} input--red`;
    },
  },
  created() {
    this.setLoading(false);
    this.locale = this.user.locale;
    this.hideDotfiles = this.user.hideDotfiles;
    this.singleClick = this.user.singleClick;
  },
  methods: {
    ...mapMutations(["updateUser", "setLoading"]),
    async updatePassword(event) {
      event.preventDefault();

      if (this.password !== this.passwordConf || this.password === "") {
        return;
      }

      try {
        const data = { id: this.user.id, password: this.password };
        await api.update(data, ["password"]);
        this.updateUser(data);
        this.$showSuccess(this.$t("settings.passwordUpdated"));
      } catch (e) {
        this.$showError(e);
      }
    },
    async updateSettings(event) {
      event.preventDefault();

      try {
        const data = {
          id: this.user.id,
          locale: this.locale,
          hideDotfiles: this.hideDotfiles,
          singleClick: this.singleClick,
        };
        await api.update(data, ["locale", "hideDotfiles", "singleClick"]);
        this.updateUser(data);
        this.$showSuccess(this.$t("settings.settingsUpdated"));
      } catch (e) {
        this.$showError(e);
      }
    },
  },
};
</script>
