<script setup>
</script>

<template>
  <h1>
    Гаджет Шоп
  </h1>
  <p>
    Приветствуем в магазине Гаджет Шоп! Здесь вы найдете самые вкусные цены на гаджеты!
  </p>
  <img src="assets/img/1.png" class="img-fluid">
    
  
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
