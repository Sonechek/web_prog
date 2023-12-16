<script setup>
</script>

<template>
  <div><button class="btn btn-primary" @click="testGet">Тест GET запроса</button></div>
  <div><button class="btn btn-primary" @click="testPost">Тест POST запроса</button></div>
  <h1>Это моя компания</h1>
  <p>Добро пожаловать!<br>Мы занимаемся тем, что создаем сайты, делаем это хорошо и качественно. Далее мы научимся создавать и полноценные веб-приложения!</p>
  <img src="/assets/img/main.jpg" class="img-fluid">
  
</template>

<script>
export default {
  data() {
    return {
    }
  },
  methods: {
    async testPost() {
      // Отправляем запрос типа POST
      const response = await fetch('/testpost', {
        method: 'POST', 
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          name: "Иван"
        }) 
      });

      const jsonResult = await response.json();

      console.log(jsonResult);
    },
    async testGet() {
      // Отправляем запрос типа GET
      const response = await fetch(`/testget?name=Василий`, {
        method: 'GET'
      });

      const textResult = await response.text();

      console.log(textResult);
    }
  }
}
</script>
