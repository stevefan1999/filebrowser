<template>
  <div id="download" class="card floating">
    <div class="card-title">
      <h2>{{ $t("prompts.download") }}</h2>
    </div>

    <div class="card-content">
      <p>{{ $t("prompts.downloadMessage") }}</p>

      <button
        v-for="(ext, format) in formats"
        :key="format"
        v-focus
        class="button button--block"
        @click="showConfirm(format)"
      >
        {{ ext }}
      </button>
    </div>
  </div>
</template>

<script>
import { mapState } from "vuex";

export default {
  name: "download",
  data: function () {
    return {
      formats: {
        zip: "zip",
        tar: "tar",
        targz: "tar.gz",
        tarbz2: "tar.bz2",
        tarxz: "tar.xz",
        tarlz4: "tar.lz4",
        tarsz: "tar.sz",
      },
    };
  },
  computed: mapState(["showConfirm"]),
};
</script>
