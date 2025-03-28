<template>
  <div class="container mt-5">
    <h1>Welcome to Home Page</h1>
    <div v-if="isLoggedIn" class="alert alert-success mt-4">
      You are logged in!
    </div>
    <div v-else class="alert alert-warning mt-4">
      Please login or register
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      isLoggedIn: !!localStorage.getItem('authToken')
    }
  }
}
</script>