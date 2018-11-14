<template>
    <div class="itemcard">
        <h1>{{ item.Name }}</h1>
        <h2>{{ item.Description }}</h2>
        <h3>{{ item.Amount }}円</h3>
    </div>
</template>

<script>
    export default{
        name: 'ItemCard',
        props: [
            'item'
        ]
    }
</script>

<style scoped>
.itemcard{
    border: solid 1px gray;
}
</style>
