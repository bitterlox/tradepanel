<template>
<!--  <head>-->
<!--    &lt;!&ndash; import reset material icons, fonts and vue-material-adapter stylesheets &ndash;&gt;-->
<!--    <link-->
<!--        rel="stylesheet"-->
<!--        href="https://fonts.googleapis.com/icon?family=Material+Icons"-->
<!--    />-->
<!--    <link-->
<!--        rel="stylesheet"-->
<!--        href="https://fonts.googleapis.com/css?family=Roboto:300,400,500"-->
<!--    />-->
<!--    <link-->
<!--        rel="stylesheet"-->
<!--        href="https://fonts.googleapis.com/css?family=Roboto+Mono:300,400,500"-->
<!--    />-->
<!--  </head>-->
  <div id=app>
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
    </div>
    <router-view/>
  </div>
</template>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
