<template>
    <div class="hello">
        <h1>{{ item.Name }}</h1>
        <h2>{{ item.Description }}</h2>
        <h3>{{ item.Amount }}円</h3>
        <payjp-checkout 
            api-key="sk_test_4a8ffddda7dbe1a3210c5615"
            text="カード情報を入力して購入"
            submit-text="購入確定"
            name-placeholder="濱　和也"
            v-on:created="onTokenCreated"
            v-on:failed="onTokenFailed">
        </payjp-checkout>

        <p>{{ message }}</p>
        <router-link to="/">HOMEへ</router-link>
    </div>
</template>

<script>
    import axios from 'axios'
    export default {
        name: 'ItemCard',
        data () {
            return {
                item: {},
                message: ''
            }
        },
        created (){
            // urlで指定された動的パラメータから商品情報をとってくる
            axios
                .get(`http://localhost:8888/api/v1/items${this.$route.param.id}`)
                .then(res => {
                    this.item = res.data
                })
        },
        beforeDestroy() {
            window.PayjpCheckout = null
        },
        methods: {
            // カードのToken化に成功したら呼ばれる。そのTokenでそのまま商品購入にうつる。
            onTokenCreated: function(res){
                console.log(res.id);
                const data = {Token: res.id}
                axios.post(`http://localhost:8888/api/v1/charge/items$(this.$route.params.id)`, data)
                .then(res => {
                    this.message = '商品の購入が完了しました！'
                })
            }
        },
        onTokenFailed: function(status, err){
            console.log(status)
            console.log(err)
        }
    }
</script>