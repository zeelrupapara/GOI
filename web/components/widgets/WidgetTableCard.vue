<template>
  <div class="mt-3">
    <div class="row">
        <!-- Table -->
        <div class="col-xl-12">
            <Portlet :headertitle="tableTitle">
              <div class="col-sm-12 col-md-6">
                  <div id="tickets-table_filter" class="dataTables_filter">
                      <label class="d-inline-flex align-items-center">
                        <b-form-select @change="onStatusChange" v-model="selectedStatus" :options="statusOptions" size="sm" class="mt-3"></b-form-select>
                      </label>
                  </div>
              </div>
              <div class="card-body pt-0">
                  <div class="table-responsive mb-0">
                      <table class="table table-hover table-centered mb-0">
                          <thead>
                              <tr>
                                  <th>Title</th>
                                  <th>Status</th>
                                  <th>Assignee</th>
                                  <th>Repository</th>
                                  <th>Organization</th>
                                  <th>Date</th>
                              </tr>
                          </thead>
                          <tbody>
                              <tr v-for="(tableDeatil, index) in tableDetails" :key="index">
                                  <td v-b-tooltip.hover :title=tableDeatil.title><a :href="tableDeatil.url" target="_blank" class="text-secondary">{{ getShortString(tableDeatil.title) }}</a></td>
                                  <td>
                                    <b-badge v-if="tableDeatil.status == 'OPEN'" variant="success">{{ tableDeatil.status }}</b-badge>
                                    <b-badge v-else-if="tableDeatil.status == 'MERGED'" variant="info">{{ tableDeatil.status }}</b-badge>
                                    <b-badge v-else-if="tableDeatil.status == 'CLOSED'" variant="danger">{{ tableDeatil.status }}</b-badge>
                                    <b-badge v-else variant="secondary">Null</b-badge>
                                  </td>
                                  <td>{{ tableDeatil.assignee}}</td>
                                  <td>{{ tableDeatil.repository }}</td>
                                  <td>{{ tableDeatil.organization }}</td>
                                  <td>{{ getFormatedDate(tableDeatil.updated_at) }}</td>
                              </tr>
                          </tbody>
                      </table>
                  </div>
                  <div class="row">
                    <div class="col">
                      <div class="float-right">
                        <b-row>
                          <b-col v-show="pageInfo.previous">
                            <button type="button" @click="previousPage" class="btn btn-sm nav-link dropdown-toggle arrow-none waves-effect waves-light"><i class="fas fa-angle-left noti-icon"></i></button>
                          </b-col>
                          <b-col v-show="pageInfo.next">
                            <button type="button" @click="nextPage" class="btn btn-sm nav-link dropdown-toggle arrow-none waves-effect waves-light"><i class="fas fa-angle-right noti-icon"></i></button>
                          </b-col>
                        </b-row>
                      </div>
                  </div>
                </div>
              </div>
            </Portlet>
        </div>
    </div>
  </div>
</template>

<script>
export default {
  props:{
    tableTitle: {
      type: String,
      default: ""
    },
    tableDetails:{
      type: Array,
      default: []
    },
    tableName:{
      type: String,
      default: ""
    },
    pageInfo:{
      type: Object,
      default: {
        previous: false,
        next: false
      }
    },
    selectedStatus:{
      type: String,
      default: null
    }
  },
  data(){
    return{
      statusOptions: []
    }
  },
  created(){
    this.setStatusFilterOptions();
  },
  watch:{
    "$route.query":{
      handler(){
        const queryParams = { ...this.$route.query };
        if (!queryParams.pr_status){
          this.selectedStatus
        }
      }
    }
  },
  methods:{
    getShortString(str){
      if (str.length > 20){
        return str.substring(0,20) + '...'
      }
      return str.substring(0,20)
    },
    getFormatedDate(date){
      return this.$utils.getFormattedTimeStamp(date)
    },
    nextPage(){
      const queryParams = { ...this.$route.query };
      switch (this.tableName) {
        case 'PR':
          if (!queryParams.pr_page) {
            queryParams.pr_page = 2;
          }else{
            queryParams.pr_page = (Number(queryParams.pr_page) + 1).toString();
          }
          break;

        case 'Issue':
          if (!queryParams.issue_page) {
            queryParams.issue_page = 2;
          }else{
            queryParams.issue_page = (Number(queryParams.issue_page) + 1).toString();
          }
          break;

        case null:
          break;
      }
      this.$router.push({
        query: queryParams
      })
    },
    previousPage(){
      const queryParams = { ...this.$route.query };
      switch (this.tableName) {
        case "PR":
          if (!queryParams.pr_page) {
            queryParams.pr_page = 1;
          }else{
            if (Number(queryParams.pr_page) > 1) {
              queryParams.pr_page = (Number(queryParams.pr_page) - 1).toString();
            }
          }
          break;
        case "Issue":
          if (!queryParams.issue_page) {
            queryParams.issue_page = 1;
          }else{
            if (Number(queryParams.issue_page) > 1) {
              queryParams.issue_page = (Number(queryParams.issue_page) - 1).toString();
            }
          }
          break;
        case null:
          break;
        }
        this.$router.push({
          query: queryParams
        })
    },
    onStatusChange(){
      const queryParams = { ...this.$route.query };
      switch (this.tableName) {
        case 'PR':
          delete queryParams.pr_status;
          if (this.selectedStatus) {
            queryParams.pr_status = this.selectedStatus;
          }
          break;

        case 'Issue':
          delete queryParams.issue_status;
          if (this.selectedStatus) {
            queryParams.issue_status = this.selectedStatus;
          }
          break;

        case null:
          break;
      }
      this.$router.push({
        query: queryParams
      })
    },
    setStatusFilterOptions(){
      if (this.tableName === 'PR') {
        this.statusOptions = [
          { value: null, text: '-- Select Status --' },
          { value: 'open', text: 'Open' },
          { value: 'merged', text: 'Merged' },
          { value: 'closed', text: 'Closed' }
        ]
      }else{
        this.statusOptions = [
          { value: null, text: '-- Select Status --' },
          { value: 'open', text: 'Open' },
          { value: 'closed', text: 'Closed' }
        ]
      }
    }
  }
}
</script>
