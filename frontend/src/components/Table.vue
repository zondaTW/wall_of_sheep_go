<template>
  <div id="tableSection" class="Table container">
    <b-table striped hover dark :items="tableData" :fields="fields"></b-table>
  </div>
</template>

<script>
export default {
  name: 'Table',
  data () {
    return {
      fields: ['IP', 'Protocol', 'Port', 'Type'],
      tableData: [
      ],
      ws: null,
    }
  },
  methods: {
    initWS: function() {
      this.ws = new WebSocket("ws://localhost:3000/ws")
      this.ws.onmessage = this.onMessage
      this.ws.onclose = this.onClose
    },
    onMessage: function (e) {
      this.tableData.unshift(JSON.parse(e.data))
      this.tableData = this.tableData.slice(0, 10)
    },
    onClose: function (e) {
      console.log("Socket closed.")
      this.ws = null
      let vm = this
      setTimeout(function () {
        vm.initWS()
      }, 5000);
    }
  },
  mounted: function() {
    this.initWS()
  }
}
</script>


<style scoped>
#tableSection {
  min-width: 300px;
  width: 70%;
  overflow-y:auto;
}
</style>
