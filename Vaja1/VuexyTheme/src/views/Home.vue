<template>
    <div>
        <b-card title="Kick start your project 🚀">
            <b-card-text>All the best for your new project.</b-card-text>
            <b-card-text>Please make sure to read our <b-link
                href="https://pixinvent.com/demo/vuexy-vuejs-admin-dashboard-template/documentation/"
                target="_blank"
            >
                Template Documentation
            </b-link> to understand where to go from here and how to use our template.</b-card-text>

            <span class="text-center">

                <b-card-text><h3>Dayjs</h3></b-card-text>
                <b-card-text><small>Instead of moment</small></b-card-text>
                <b-card-text>Usage: {{ $dayjs().format("HH:mm:ss DD.MM.YYY") }}</b-card-text>
            </span>
        </b-card>

        <b-card title="Feather icons">
            <b-card-text class="text-center">
                <feather-icon
                    icon="SunIcon"
                    size="16"
                />
                Usage of feather icons
                <feather-icon
                    icon="BellIcon"
                    size="16"
                />
            </b-card-text>
            <b-card-text>Feather icons je potrebno dodati vsako posebej, ker na tak način prihranimo na build size. Če importamo vse ikone, importamo 375kb ikon, drugače pa samo 22kb</b-card-text>
            <b-card-text>Ikone importaš v <b>src\@core\components\feather-icon\FeatherIcon.vue</b></b-card-text>
            <p>
                import FileIcon from 'vue-feather-icons/icons/FileIcon'<br>
                const icons = { FileIcon }
            </p>
        </b-card>

        <b-card title="Font awesome">
            <b-card-text class="text-center">
                <fa
                    class="text-danger"
                    size="2x"
                    icon="air-freshener"
                />
                Usage of font-awesome-icons
                <fa
                    class="text-success"
                    size="2x"
                    icon="ambulance"
                />
            </b-card-text>
            <b-card-text>FontAwesome ikone je potrebno dodati vsako posebej, saj na tak način prihranimo s prostorom. Če importamo vse ikone, importamo 800kb ikon, če pa vsako posebaj, pa samo 95kb.</b-card-text>
            <b-card-text>Ikone importaš v <b>/libs/font-awesome.js</b></b-card-text>
            <p>
                import { faAirFreshener } from '@fortawesome/free-solid-svg-icons/faAirFreshener'<br>
                library.add(faAirFreshener)
            </p>
        </b-card>

        <b-card title="Tailwind CSS">
            <h1 class="text-3xl font-bold underline text-center">
                TailwindCSS is working!
            </h1>

            <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-5 xl:grid-cols-auto gap-4 text-white font-bold rounded-lg text-center">
                <div class="p-1 w-full h-14 bg-blue-500 shadow-lg rounded-lg">A</div>
                <div class="p-1 w-full h-14 bg-blue-500 shadow-lg rounded-lg">B</div>
                <div class="p-1 w-full h-14 bg-blue-500 shadow-lg rounded-lg">C</div>
                <div class="p-1 w-full h-14 bg-blue-500 shadow-lg rounded-lg">D</div>
                <div class="p-1 w-full h-14 bg-blue-500 shadow-lg rounded-lg">E</div>
                <div class="p-1 w-full h-14 bg-blue-500 shadow-lg rounded-lg">F</div>
            </div>
        </b-card>

      <b-card title="Vue select">
        <vue-select :options="rows" label="name" />
      </b-card>

      <b-card title="i18n">
        {{ $t('general.greeting') }}
        <b-button @click="onChangeLocale">Change Locale</b-button>
      </b-card>

      <b-card title="Vue good table">
        <vue-good-table
            :columns="columns"
            :rows="rows"/>
      </b-card>

        <b-card title="Want to integrate JWT? 🔒">
            <b-card-text>We carefully crafted JWT flow so you can implement JWT with ease and with minimum efforts.</b-card-text>
            <b-card-text>Please read our  JWT Documentation to get more out of JWT authentication.</b-card-text>
        </b-card>
    </div>
</template>

<script>
    import { BCard, BCardText, BLink, BButton } from 'bootstrap-vue'
    import VueSelect from 'vue-select'
    import 'vue-select/dist/vue-select.css';
    import 'vue-good-table/dist/vue-good-table.css'
    import { VueGoodTable } from 'vue-good-table';
    import i18n from '@/libs/i18n'

    export default {
        components: {
            BCard,
            BCardText,
            BLink,
          BButton,
          VueSelect,
          VueGoodTable
        },
      data() {
          return {
            currentLocale: 'en',
            columns: [
              {
                label: 'Name',
                field: 'name',
              },
              {
                label: 'Age',
                field: 'age',
                type: 'number',
              },
              {
                label: 'Percent',
                field: 'score',
                type: 'percentage',
              },
            ],
            rows: [
              { id:1, name:"John", age: 20,score: 0.03343 },
              { id:2, name:"Jane", age: 24, score: 0.03343 },
              { id:3, name:"Susan", age: 16, score: 0.03343 },
              { id:4, name:"Chris", age: 55, score: 0.03343 },
              { id:5, name:"Dan", age: 40, score: 0.03343 },
              { id:6, name:"John", age: 20, score: 0.03343 },
            ],
          }
      },
      methods: {
        onChangeLocale() {
          this.currentLocale = this.currentLocale === 'en' ? 'hr' : 'en'
          i18n.locale = this.currentLocale
        }
      }
    }
</script>
<style>

</style>
