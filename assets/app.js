// post new tweet
 var login_form = new Vue({
     el: '#input',
     data:{
         nil : false,
         user: {},
     },
     methods: {
         login: function() { // undefined
             //console.log(this.tweet.body == undefined )
                 console.log("no data");
 
                 // hide wornning message
                 // this.nil = false,
 
                 // post tweeti
                 axios.post('http://localhost:8080/login',
                 this.user,
                 // or post : {title: "new title", body: "new body"}
                 )
                 .then(function (response) {
                     console.log(response.data)
                 })
                 .catch(function (error) {
                      console.log("Have an: "+ error)
                 });
             } 
         },
})
/*
Vue.component('test', {
    template:'<h1>test component </h1>'
})

new Vue({
    el:'#test'
})

auth = new Vue({
   el: '#auth',
   data: {
       data: []
   },
   //mounted() {}
 });
 
 // get all tweets
 var Get_tweets = new Vue({
     el: '#get_tweets',
     
     methods: {
        get_tweets: function() {
             axios.get('http://localhost:8000/tweets')
             .then(response => (tweets.tweets = response.data))
        }
     }  
 })
// delete tweet
var delete_tweet = new Vue({
    el : '#del_tweet',
    data: {
        id: 20,
    },
    methods:{
        delete_tweet: function() { // undefined
            // delete tweeti
            axios.delete('http://localhost:8000/tweet/'+this.id)
            //.then(function (response) {
                //console.log(response.data)
            //})
            .catch(function (error) {
                 console.log("Have an: "+ error)
            });
        } 
    }
})

 */
