<template>
  <div class="home">
<!--    <img @click="getMessage" alt="Vue logo" src="../assets/appicon.png" :style="{ height: '400px' }"/>-->
    <input v-model="param">
    <button @click="getMessage(this.param)" >
      <div style="height: 20px; width: 20px;"></div>
    </button>
    <HelloWorld :msg="response" />
  </div>
</template>

<script lang="ts">
import { ref, defineComponent } from "vue";
import HelloWorld from "@/components/HelloWorld.vue"; // @ is an alias to /src

interface Methods {
  Greet(arg0: string): Promise<string>;
}

interface Backend {
  Server: Methods
}

declare global {
  interface Window {
    backend: Backend;
  }
}
export default defineComponent({
  name: "Home",
  components: {
    HelloWorld,
  },
  data () {
    return {
      param: ""
    }
  },
  setup() {

    const message = ref("Click the Icon");

    const getMessage = (send: string) => {
      window.backend.Server.Greet(send).then(result => {
        message.value = result;
      });
    }

    return { response: message, getMessage: getMessage };
  },
});
</script>
