// var Vue = require('vue');

Vue.component('event-item', {
  props: ['event', 'index'],
  template:'\
  <span>\
        <h4 class="list-group-item-heading"><i class="glyphicon glyphicon-bullhorn"></i>  {{event.title }}</h4>\
        <h5><i class="glyphicon glyphicon-calendar" v-if="event.date"></i> {{ event.date }}</h5>\
        <p class="list-group-item-text" v-if="event.detail">{{ event.detail }}</p>\
        <button class="btn btn-xs btn-danger" v-on:click="deleteEvent(index)">Delete</button> </span>',
  methods: {  
     deleteEvent: function (index) {
      if (confirm('Really want to delete？')) {
        console.log(index);
        this.$http.delete('/api/events/' + index)
          .then(response => response)
          .then( result =>{
              console.log(result);
              app.events.splice(index, 1);
            }).catch( err => {
              console.log(err);
              alert("unable to delete")
            });
      }
    }
  }
});

var app = new Vue({
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
            this.event = { title: '', detail: '', date: '' };
          }).catch( err => {
            console.log(err);
          });
      }

    }
  }
});
