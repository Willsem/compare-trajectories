import $ from 'jquery';
import config from '../configs/client.json';

export function filtering(trajectory) {
  let filteredTrajectory;

  $.ajax({
    url: config.server + '/filter',
    method: 'post',
    async: false,
    data: JSON.stringify(trajectory),
  })
  .done(function(result) {
    filteredTrajectory = result;
  });

  return filteredTrajectory;
}
