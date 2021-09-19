<template>
  <div>
    <!-- <title-bar :title-stack="titleStack"/> -->
    <!-- <hero-bar>
      Forms
      <router-link slot="right" to="/" class="button">
        Dashboard
      </router-link>
    </hero-bar> -->
    <section class="section is-main-section">
      <card-component title="Forms" icon="ballot">
        <form @submit.prevent="submit">
          <b-field label="IndexService" horizontal>
            <b-field :addons="false" class="field has-addons" >
              <p class="control">
                <a class="button is-static">
                  psm:
                </a>
              </p>
              <b-input icon="" type="text" v-model="form.psm" placeholder="psm" name="psm" required expanded />
            </b-field>
            <b-field  :addons="false" class="field has-addons">
              <p class="control">
                <a class="button is-static">
                  index_name:
                </a>
              </p>
              <b-input icon="" type="text" v-model="form.index_name" placeholder="index_name" name="index_name" required expanded/>
            </b-field>
          </b-field>
          <b-field message="Do not enter the leading zero" horizontal>
            <b-field>
              <p class="control">
                <a class="button is-static">
                  key
                </a>
              </p>
              <b-input type="text" v-model="form.key" name="key" expanded />
            </b-field>
          </b-field>
          <b-field label="pb" horizontal>
            <b-select placeholder="Select a pb" v-model="form.pb" required>
              <option v-for="(pb, index) in pb" :key="index" :value="pb">
                {{ pb }}
              </option>
            </b-select>
          </b-field>
          <b-field label="Switch" horizontal>
          <b-switch v-model="form.is_pack"  type="is-info">
            is_pack
          </b-switch>
          </b-field>
          <hr>
          <b-field label="Radio" class="has-check" horizontal>
          <radio-picker :options="{one:'123',two:'12322'}" v-model="form.radio"></radio-picker>
          </b-field>
          <hr>
          <b-field horizontal>
            <b-field grouped>
              <div class="control">
                <b-button native-type="submit" type="is-primary">Submit</b-button>
              </div>
              <div class="control">
                <b-button type="is-primary is-outlined" @click="reset">Reset</b-button>
              </div>
            </b-field>
          </b-field>
        </form>
      </card-component>
      <card-component title="test" icon="ballot">
        <div>
        <v-text-field
          label="Main input"
          :rules="rules"
          v-model="form.psm"
          hide-details="auto"
        ></v-text-field>
        <v-text-field label="Another input"></v-text-field>
      </div>
      </card-component>
    </section>
  </div>
</template>

<script>
// import TitleBar from '@/components/TitleBar'
import CardComponent from '@/components/CardComponent'
import mapValues from 'lodash/mapValues'
// import CheckboxPicker from '@/components/CheckboxPicker'
import RadioPicker from '@/components/RadioPicker'
// import FilePicker from '@/components/FilePicker'
// import HeroBar from '@/components/HeroBar'

export default {
  name: 'Forms',
  components: { RadioPicker, CardComponent },
  data () {
    return {
      isLoading: false,
      form: {
        psm: null,
        index_name: null,
        key: null,
        pb: null,
        radio: null,
        is_pack: true
      },
      pb: [
        'example',
        'ecom',
        'Sales'
      ],
      rules: [
        value => !!value || 'Required.'
      ]
    }
  },
  computed: {
  },
  methods: {
    submit () {
      console.log(JSON.stringify(this.form))
      this.$http.post('api/x', this.form).then((res) => {
        alert(JSON.stringify(res.data.data))
      })
    },
    reset () {
      this.form = mapValues(this.form, item => {
        if (item && typeof item === 'object') {
          return []
        }
        return null
      })

      this.$buefy.snackbar.open({
        message: 'Reset successfully',
        queue: false
      })
    }
  }
}
</script>
