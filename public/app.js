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

      this.$http.get('/api/events').then((response) => {
        function a(events){
          // success callback
         Vue.set('events', events);
          console.log("success in getting events")  
        }
        
      }, (response) => {
        // error callback
        function failfetch (err) {
          console.log(err);
        }
      });
        
    },

    addEvent: function () {
      if (this.event.title.trim()) {
        // this.events.push(this.event);
        // this.event = { title: '', detail: '', date: '' };
        this.$http.post('/api/events', this.event)
          .then(function (res) {
            this.events.push(this.event);
            console.log('Event added!');
          },function (err) {
            console.log(err);
          });
      }
    },

    deleteEvent: function (index) {
      if (confirm('Really want to delete？')) {
        // this.events.splice(index, 1);
        this.$http.delete('/api/events/' + event.id)
          .then(function (res) {
            console.log(res);
            this.events.splice(index, 1);
          },
          function (err) {
            console.log(err);
          });
      }
    }
  }
});
