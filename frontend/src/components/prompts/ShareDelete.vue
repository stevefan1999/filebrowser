<template>
  <div class="card floating">
    <div class="card-content">
      <p>{{ $t("prompts.deleteMessageShare", { path: "" }) }}</p>
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
        :aria-label="$t('buttons.delete')"
        :title="$t('buttons.delete')"
        class="button button--flat button--red"
        @click="submit"
      >
        {{ $t("buttons.delete") }}
      </button>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";

export default {
  name: "share-delete",
  computed: {
    ...mapState(["showConfirm"]),
  },
  methods: {
    submit: function () {
      this.showConfirm();
    },
  },
};
</script>
