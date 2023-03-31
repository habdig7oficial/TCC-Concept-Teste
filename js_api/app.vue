<template>
  <!-- you will need to handle a loading state -->
  <div v-if="pending">
    Loading ...
  </div>
  <div v-else>
    <div>
      ok - {{posts}}
    </div>

    <h1>licensa: {{  posts.license }}</h1>
    <h1>Query Params: {{  posts.query_params }}</h1>
    <h1>Query Params: {{  posts.dbvalue }}</h1>
  </div>
</template>

<script setup lang="ts">
const { pending, data: posts }:any = useLazyFetch('http://0.0.0.0:7777/')
watch(posts, (newPosts) => {
  // Because posts starts out null, you will not have access
  // to its contents immediately, but you can watch it.
})

console.log(posts)
</script>
