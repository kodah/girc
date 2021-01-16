<template>
  <v-app id="inspire">
    <v-app-bar
        app
        clipped-right
        flat
        height="50"
    >
      <v-spacer></v-spacer>

      <v-responsive max-width="156">
        <v-text-field
            dense
            flat
            height="35"
            hide-details
            rounded
            solo-inverted
        ></v-text-field>
      </v-responsive>
    </v-app-bar>

    <v-navigation-drawer
        v-model="drawer"
        app
        width="300"
    >
      <v-navigation-drawer
          v-model="drawer"
          absolute
          color="grey darken-3"
          mini-variant
      >
        <v-avatar
            class="d-block text-center mx-auto mt-4"
            color="primary darken-1"
            size="36"
        >
          <v-icon>mdi-forum</v-icon>
        </v-avatar>

        <v-divider class="mx-3 my-5"></v-divider>

        <v-avatar
            v-for="n in 6"
            :key="n"
            class="d-block text-center mx-auto mb-9"
            color="grey lighten-1"
            size="28"
        ></v-avatar>
      </v-navigation-drawer>

      <UserStatus user="kodah" network="chat.freenode.org:6697"></UserStatus>

      <ChannelList></ChannelList>
    </v-navigation-drawer>

    <v-navigation-drawer
        app
        clipped
        right
    >
      <v-list>
        <v-list-item
            link
        >
          <v-list-item-content>
            <v-list-item-title class="red--text font-weight-bold text-body-1">kodah</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-main>
      <Message me=true user="kodah" message="waddup yall" time="Today at 10:41 AM"></Message>
    </v-main>

    <MessageBox></MessageBox>
  </v-app>
</template>

<script>
import MessageBox from "@/components/MessageBox";
import Message from "@/components/Message";
import UserStatus from "@/components/UserStatus";
import Pane from "@/components/ChannelList";
import ChannelList from "@/components/ChannelList";
export default {
  name: 'App',
  components: {ChannelList, Pane, UserStatus, Message, MessageBox },
  data: () => ({ drawer: null }),
  methods: {}
}
</script>
