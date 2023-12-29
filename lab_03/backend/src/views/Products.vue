<script setup>
</script>

<style>
.product-card {
  border: 1 px solid gray;
  border-radius: 1rem;
  padding: 1rem;
  display: flex;
  align-items: center;
  margin-top: 1rem;
}
.product-image {
  max-width: 100px;
}
.product-text {
  padding: 1rem;
}
.product-name {
  font-size: 24px;
  font-weight: bold;
}
</style>

<template>
  <nav aria-label="Page navigation example">
    <ul class="pagination">
      <li class="page-item" @click="minusPage()"><a class="page-link" href="#">Предыдущая</a></li>
      <li :class="{ 'page-item':true, 'active': page == current_page}" v-for="page in pageCount">
        <a class="page-link " href="#" @click="current_page = page">{{ page }}</a>
      </li>
      <li class="page-item" @click="plusPage()"><a class="page-link" href="#">Следующая</a></li>
    </ul>
  </nav>
  <div v-for="product in products" class="product-card">
    <div>
      <img :src="product.img_url" class="product-image">
    </div>
  <div class="product-text">
    <div class="product-name">
      {{ product.name }}
    </div>
    <div>
      {{ product.description }}
    </div>
    <div>
      {{ product.price }}
    </div>
  </div>
  </div>
</template> 

<script>
export default {
  data() {
    return {
      products: [],
      products_per_page: 0,
      product_count: 0,
      current_page: 1,
    }
  },
  watch: {
    current_page () {
        this.getProducts()
    }
  },
  computed: { 
    pageCount() {
      if (this.products_per_page == 0){
        return 0;
      }
      return Math.ceil(this.product_count / this.products_per_page);
    }
  },
  methods: {
    minusPage(){
      this.current_page--;
      if (this.current_page < 1){
        this.current_page = 1;
      }
    },
    plusPage(){
      this.current_page++;
      if (this.current_page > this.pageCount){
        this.current_page = this.pageCount;
      }
    },
    async getProducts() {
      const response = await fetch('/getproducts', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          page: this.current_page
        })
      });

      const jsonResult = await response.json();
      this.products = jsonResult;

      console.log(jsonResult);
    },
    async getProductCount() {
      const response = await fetch('/getproductcount', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        
      });

      const jsonResult = await response.json();
      this.product_count = jsonResult.count;
      this.products_per_page = jsonResult.products_per_page;
    },
  },
  mounted() {
    this.getProductCount();
    this.getProducts();
    
  }
}
</script>
