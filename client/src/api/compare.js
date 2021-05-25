import $ from 'jquery';
import config from '../configs/client.json';

export function compare(perfectTrajectory, compareTrajectory) {
  let compareResult;

  $.ajax({
    url: config.server + '/compare',
    method: 'post',
    async: false,
    data: JSON.stringify({
      "perfect": perfectTrajectory,
      "compared": compareTrajectory,
    }),
  })
  .done(function(result) {
    compareResult = result;
  });

  return compareResult;
}
