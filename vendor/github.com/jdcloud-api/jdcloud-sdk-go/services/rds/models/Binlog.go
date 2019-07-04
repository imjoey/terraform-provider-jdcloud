// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package models


type Binlog struct {

    /* binlog日志备份ID (Optional) */
    BinlogBackupId string `json:"binlogBackupId"`

    /* binlog日志名称 (Optional) */
    BinlogName string `json:"binlogName"`

    /* binlog日志大小，单位KB (Optional) */
    BinlogSizeKB int64 `json:"binlogSizeKB"`

    /* binlog开始时间,格式为：YYYY-MM-DD HH:mm:ss (Optional) */
    BinlogStartTime string `json:"binlogStartTime"`

    /* binlog结束时间,格式为：YYYY-MM-DD HH:mm:ss (Optional) */
    BinlogEndTime string `json:"binlogEndTime"`
}
