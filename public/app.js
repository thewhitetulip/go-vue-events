// var Vue = require('vue');

new Vue({
  el: '#events',

  data: {
    event: { title: '', detail: '', date: '' },
    events: []  
  },
  delimiters: ['${', '}'],
  

  mounted: function () {
    this.fetchEvents();
  },

  methods: {

    fetchEvents: function () {
        var events = [];
        this.$http.get('/api/events/')
        .then(response => response.json())
        .then(result => {
           Vue.set(this.$data, 'events', result);
            console.log("success in getting events")  
        })
        .catch(err => {
            console.log(err);
        });
    },

    addEvent: function () {
      if (this.event.title.trim()) {
        // this.events.push(this.event);
        // this.event = { title: '', detail: '', date: '' };
        this.$http.post('/api/events/', this.event,{emulateJSON: true})
          .then(response => response)
          .then( result => {
            this.events.push(this.event);
            console.log('Event added!');
          }).catch( err => {
            console.log(err);
          });
      }

    },

    deleteEvent: function (index) {
      if (confirm('Really want to delete？')) {
        console.log(index);
        this.$http.delete('/api/events/' + index)
          .then(response => response)
          .then( result =>{
              console.log(result);
              this.events.splice(index, 1);
            }).catch( err => {
              console.log(err);
              alert("unable to delete")
            });
      }
    }
  }
});
