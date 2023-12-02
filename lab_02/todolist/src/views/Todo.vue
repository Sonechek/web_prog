<script setup>
// Эта часть выполняется при первоначальной загрузке компонента
</script>
<template>

<div>
    <h2>Add Note</h2>
    <form @submit.prevent="createNote">
      <input type="text" v-model="title" placeholder="Note Title" />
      <br />
      <textarea v-model="content" placeholder="Note Content"></textarea>
      <br />
      <button type="submit">Create Note</button>
    </form>
  </div>
</template>
<script>
// Здесь исходный код компонента
</script>
