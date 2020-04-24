<template>
    <div id="main" v-cloak v-on:click="hideTooltip">

        <div class="tooltip" v-on:click.stop v-if="show_tooltip">

            <input type="text" v-model="text_content"/>
        </div>

        <p v-on:click.stop="toggleTooltip">{{text_content}}</p>

        <button v-on:click="cmdb">跨域</button>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                show_tooltip: false,
                text_content: 'Edit me.'
            }
        },
        methods: {
            hideTooltip: function(){
                // When a model is changed, the view will be automatically updated.
                this.show_tooltip = false;
            },
            toggleTooltip: function(){
                this.show_tooltip = !this.show_tooltip;
            },
            cmdb:function () {
                console.log('ready...');

                this.axios.get('/app.json')
                    .then(response => (console.log(response)));
            }
        },
        created: function(){
            this.toggleTooltip()
        }

    }
</script>