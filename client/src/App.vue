<template>
    <div id="app">
        <router-view />
    </div>
</template>

<script>
export default {
    created() {
        this.$axios.defaults.headers.common['Authorization'] = localStorage.getItem("token")
    }
}
</script>

<style lang="scss">
@import "~bulma/sass/utilities/_all";

$primary: #00d1b2;
$primary-invert: findColorInvert($primary);

$link: $primary;
$link-invert: $primary-invert;
$link-focus-border: $primary;

@import "~bulma";
@import '~buefy/src/scss/buefy';

</style>

<style>
#app {
	padding: 20px;
}
</style>
