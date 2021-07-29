<template>
  <div class="card floating">
    <div class="card-title">
      <h2>{{ $t("prompts.newFile") }}</h2>
    </div>

    <div class="card-content">
      <p>{{ $t("prompts.newFileMessage") }}</p>
      <input
        v-focus
        v-model.trim="name"
        class="input input--block"
        type="text"
        @keyup.enter="submit"
      />
    </div>

    <div class="card-action">
      <button
        :aria-label="$t('buttons.cancel')"
        :title="$t('buttons.cancel')"
        class="button button--flat button--grey"
        @click="$store.commit('closeHovers')"
      >
        {{ $t("buttons.cancel") }}
      </button>
      <button
        :aria-label="$t('buttons.create')"
        :title="$t('buttons.create')"
        class="button button--flat"
        @click="submit"
      >
        {{ $t("buttons.create") }}
      </button>
    </div>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import { files as api } from "@/api";
import url from "@/utils/url";

export default {
  name: "new-file",
  data: function () {
    return {
      name: "",
    };
  },
  computed: {
    ...mapGetters(["isFiles", "isListing"]),
  },
  methods: {
    submit: async function (event) {
      event.preventDefault();
      if (this.new === "") return;

      // Build the path of the new directory.
      let uri = this.isFiles ? this.$route.path + "/" : "/";

      if (!this.isListing) {
        uri = url.removeLastDir(uri) + "/";
      }

      uri += encodeURIComponent(this.name);
      uri = uri.replace("//", "/");

      try {
        await api.post(uri);
        this.$router.push({ path: uri });
      } catch (e) {
        this.$showError(e);
      }

      this.$store.commit("closeHovers");
    },
  },
};
</script>
