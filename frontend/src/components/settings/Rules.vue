<template>
  <form class="rules small">
    <div v-for="(rule, index) in rules" :key="index">
      <input v-model="rule.regex" type="checkbox" /><label>Regex</label>
      <input v-model="rule.allow" type="checkbox" /><label>Allow</label>

      <input
        v-if="rule.regex"
        v-model="rule.regexp.raw"
        :placeholder="$t('settings.insertRegex')"
        type="text"
        @keypress.enter.prevent
      />
      <input
        v-else
        v-model="rule.path"
        :placeholder="$t('settings.insertPath')"
        type="text"
        @keypress.enter.prevent
      />

      <button class="button button--red" @click="remove($event, index)">
        -
      </button>
    </div>

    <div>
      <button class="button" default="false" @click="create">
        {{ $t("buttons.new") }}
      </button>
    </div>
  </form>
</template>

<script>
export default {
  name: "rules-textarea",
  props: ["rules"],
  methods: {
    remove(event, index) {
      event.preventDefault();
      let rules = [...this.rules];
      rules.splice(index, 1);
      this.$emit("update:rules", [...rules]);
    },
    create(event) {
      event.preventDefault();

      this.$emit("update:rules", [
        ...this.rules,
        {
          allow: true,
          path: "",
          regex: false,
          regexp: {
            raw: "",
          },
        },
      ]);
    },
  },
};
</script>
