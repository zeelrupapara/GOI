<script>
import FullCalendar from '@fullcalendar/vue'
import dayGridPlugin from '@fullcalendar/daygrid'
import timeGridPlugin from '@fullcalendar/timegrid'
import interactionPlugin from '@fullcalendar/interaction'
import bootstrapPlugin from '@fullcalendar/bootstrap'
import listPlugin from '@fullcalendar/list'

export default {
  components: {
    FullCalendar,
  },
  head() {
    return {
      title: `${this.title} | GPAT`,
    }
  },
  data() {
    return {
      title: 'Sync',
      items:[
        {
          text: 'Home',
          to: '/',
        },
        {
          text: 'Sync',
          active: true,
        },
      ],
      isLoading: false,
      calendarOptions: {
        headerToolbar: {
          left: 'prev,next today',
          center: 'title',
          right: 'dayGridMonth,timeGridWeek,timeGridDay,listWeek',
        },
        plugins: [
          dayGridPlugin,
          timeGridPlugin,
          interactionPlugin,
          bootstrapPlugin,
          listPlugin,
        ],
        initialView: 'dayGridMonth',
        themeSystem: 'bootstrap',
        initialEvents: {},
        editable: true,
        droppable: true,
        eventResizableFromStart: true,
        eventsSet: this.handleEvents,
        weekends: true,
        selectable: true,
        selectMirror: true,
        dayMaxEvents: true,
      },
      event: {
        title: '',
        category: '',
      },
    }
  },
  created() {
    this.getSyncedDates()
  },
  methods: {
    getTimestampFromISODate(timeStr) {
      const date = new Date(timeStr)
      return date.getTime()
    },
    getSyncedDates() {
      this.isLoading = true
      this.$axios
        .get(`${this.$constants.API_URL_PREFIX}/sync`)
        .then((res) => {
          if (res.data.data.length > 0) {
            const syncedData = res.data.data
            let id = 0
            this.calendarOptions.initialEvents = syncedData.map((event) => {
              id++
              return {
                id: id,
                start: this.getTimestampFromISODate(event.unique_date),
                title: event.activities.join(', '),
                className: 'bg-primary text-white',
              }
            })
          }
        })
        .catch((err) => {
          this.$toaster.error(err)
        })
        .finally(() => {
          this.isLoading = false
        })
    },
  },
}
</script>

<template>
  <div>
    <PageHeader :items="items" />
    <Loading v-if="isLoading" />
    <div v-else class="row">
      <div class="col-12">
        <div class="card">
          <div class="card-body">
            <div class="app-calendar">
              <FullCalendar
                ref="fullCalendar"
                :options="calendarOptions"
              ></FullCalendar>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
