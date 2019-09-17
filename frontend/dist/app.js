

var demo = new Vue({

    el: '#main',

    data: {
        show_tooltip: false,
        text_content: 'Edit me.'
    },

    // Functions we will be using.
    methods: {
        hideTooltip: function(){
            // When a model is changed, the view will be automatically updated.
            this.show_tooltip = false;
        },
        toggleTooltip: function(){
            this.show_tooltip = !this.show_tooltip;
        }
    }
    });