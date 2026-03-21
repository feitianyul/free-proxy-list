**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-03-21 07:43:23 UTC（2026-03-21 15:43:23 UTC+8）

**代理总数：215**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 215 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 215 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 38.34.179.83:8448 | ✓ 602ms | ✓ 1061ms | ✓ 946ms | ✓ 1100ms | ✓ 817ms | http |
| 38.34.179.105:8449 | ✓ 599ms | ✓ 879ms | ✓ 963ms | ✓ 1277ms | ✓ 811ms | http |
| 147.161.239.240:8800 | ✓ 1027ms | 否 | ✓ 1611ms | ✓ 1722ms | ✓ 1346ms | http |
| 38.34.179.61:8445 | ✓ 1145ms | ✓ 1854ms | ✓ 783ms | ✓ 1603ms | ✓ 668ms | http |
| 147.161.210.140:8800 | 否 | 否 | ✓ 1021ms | ✓ 1110ms | ✓ 1335ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1348ms | ✓ 1443ms | ✓ 1498ms | http |
| 137.220.150.152:6005 | 否 | 否 | ✓ 1547ms | ✓ 1705ms | ✓ 1389ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1767ms | ✓ 1513ms | ✓ 1628ms | http |
| 167.103.34.108:8800 | 否 | 否 | ✓ 1676ms | ✓ 1778ms | ✓ 1806ms | http |
| 45.167.124.52:8080 | ✓ 797ms | 否 | ✓ 1388ms | ✓ 1878ms | ✓ 1663ms | http |
| 38.34.183.130:8452 | ✓ 332ms | ✓ 1003ms | ✓ 241ms | ✓ 970ms | ✓ 647ms | http |
| 38.34.179.104:8446 | ✓ 538ms | ✓ 1144ms | ✓ 229ms | ✓ 890ms | ✓ 713ms | http |
| 38.34.179.98:8453 | ✓ 877ms | ✓ 878ms | ✓ 274ms | ✓ 889ms | ✓ 666ms | http |
| 38.34.178.186:8451 | ✓ 305ms | ✓ 1361ms | ✓ 300ms | ✓ 917ms | ✓ 724ms | http |
| 38.34.179.25:8444 | ✓ 337ms | ✓ 958ms | ✓ 262ms | ✓ 1119ms | ✓ 908ms | http |
| 38.34.179.23:8444 | ✓ 378ms | 否 | ✓ 297ms | ✓ 939ms | ✓ 714ms | http |
| 38.34.179.186:8444 | ✓ 340ms | ✓ 1568ms | ✓ 251ms | ✓ 876ms | ✓ 1036ms | http |
| 45.136.130.171:8445 | ✓ 276ms | ✓ 1534ms | ✓ 875ms | ✓ 887ms | ✓ 788ms | http |
| 38.34.179.101:8446 | ✓ 544ms | 否 | ✓ 251ms | ✓ 919ms | ✓ 694ms | http |
| 38.34.179.96:8451 | ✓ 1234ms | ✓ 901ms | ✓ 256ms | ✓ 890ms | ✓ 947ms | http |
| 38.145.208.185:8449 | ✓ 329ms | ✓ 1930ms | ✓ 875ms | ✓ 897ms | ✓ 699ms | http |
| 38.34.179.88:8446 | ✓ 953ms | ✓ 1784ms | ✓ 327ms | ✓ 939ms | ✓ 851ms | http |
| 38.34.179.89:8444 | ✓ 814ms | ✓ 1679ms | ✓ 350ms | ✓ 1074ms | ✓ 1043ms | http |
| 38.34.179.91:8444 | ✓ 799ms | ✓ 1846ms | ✓ 356ms | ✓ 1057ms | ✓ 1084ms | http |
| 38.34.179.85:8444 | ✓ 809ms | 否 | ✓ 325ms | ✓ 1131ms | ✓ 1185ms | http |
| 38.145.203.19:8449 | ✓ 855ms | 否 | ✓ 292ms | ✓ 895ms | ✓ 705ms | http |
| 38.34.178.245:8446 | ✓ 1264ms | ✓ 1810ms | ✓ 324ms | ✓ 918ms | ✓ 727ms | http |
| 38.145.208.241:8453 | ✓ 794ms | ✓ 1231ms | ✓ 343ms | ✓ 882ms | ✓ 752ms | http |
| 38.145.208.244:8448 | ✓ 484ms | ✓ 1041ms | ✓ 853ms | ✓ 986ms | ✓ 856ms | http |
| 45.136.131.62:8449 | ✓ 431ms | 否 | ✓ 886ms | ✓ 1194ms | ✓ 827ms | http |
| 38.145.208.242:8451 | ✓ 1171ms | ✓ 1509ms | ✓ 304ms | ✓ 980ms | ✓ 1010ms | http |
| 38.34.179.39:8452 | ✓ 851ms | ✓ 1217ms | ✓ 934ms | ✓ 872ms | ✓ 808ms | http |
| 45.136.131.53:8452 | ✓ 873ms | 否 | ✓ 1057ms | ✓ 919ms | ✓ 872ms | http |
| 38.34.179.17:8446 | ✓ 1393ms | ✓ 1619ms | ✓ 467ms | ✓ 1151ms | ✓ 1317ms | http |
| 38.34.179.13:8446 | ✓ 1318ms | ✓ 1712ms | ✓ 533ms | ✓ 1192ms | ✓ 1494ms | http |
| 45.136.130.178:8449 | ✓ 298ms | ✓ 1385ms | ✓ 1014ms | ✓ 1328ms | ✓ 659ms | http |
| 38.34.179.57:8453 | ✓ 990ms | ✓ 825ms | ✓ 337ms | ✓ 1383ms | ✓ 1744ms | http |
| 38.34.179.172:8451 | ✓ 252ms | ✓ 1611ms | ✓ 504ms | ✓ 1341ms | ✓ 691ms | http |
| 38.34.179.165:8446 | ✓ 281ms | ✓ 1140ms | ✓ 386ms | ✓ 1329ms | ✓ 812ms | http |
| 137.220.151.110:6005 | ✓ 899ms | 否 | ✓ 878ms | ✓ 1234ms | ✓ 1019ms | http |
| 137.220.150.22:6005 | ✓ 1132ms | 否 | ✓ 852ms | ✓ 1221ms | ✓ 1023ms | http |
| 38.34.179.86:8452 | ✓ 1252ms | 否 | ✓ 566ms | ✓ 1408ms | ✓ 917ms | http |
| 38.34.179.203:8451 | ✓ 331ms | 否 | ✓ 1730ms | ✓ 1003ms | ✓ 1769ms | http |
| 38.34.179.54:8446 | ✓ 893ms | 否 | ✓ 567ms | ✓ 1492ms | ✓ 682ms | http |
| 38.34.179.51:8449 | ✓ 972ms | ✓ 1412ms | ✓ 427ms | ✓ 1429ms | ✓ 1146ms | http |
| 38.34.179.48:8449 | ✓ 978ms | ✓ 1360ms | ✓ 494ms | ✓ 1494ms | ✓ 1041ms | http |
| 38.34.179.97:8448 | ✓ 1637ms | ✓ 1926ms | ✓ 1093ms | ✓ 1324ms | ✓ 659ms | http |
| 137.220.150.170:6005 | ✓ 943ms | 否 | ✓ 1161ms | ✓ 1431ms | 否 | http |
| 38.34.179.19:8449 | ✓ 1580ms | ✓ 1703ms | ✓ 1119ms | 否 | ✓ 1004ms | http |
| 38.34.179.8:8449 | ✓ 1557ms | 否 | ✓ 1583ms | 否 | ✓ 1060ms | http |
| 38.34.179.11:8449 | ✓ 1586ms | 否 | ✓ 1504ms | 否 | ✓ 981ms | http |
| 45.136.130.169:8444 | ✓ 1377ms | ✓ 1642ms | ✓ 465ms | ✓ 1578ms | ✓ 1895ms | http |
| 45.88.0.114:3128 | 否 | 否 | ✓ 748ms | ✓ 1321ms | ✓ 1132ms | http |
| 38.34.179.228:8453 | ✓ 601ms | ✓ 1238ms | 否 | ✓ 1396ms | 否 | http |
| 45.136.130.167:8444 | ✓ 1386ms | 否 | ✓ 798ms | 否 | ✓ 1152ms | http |
| 213.220.62.62:3128 | 否 | 否 | ✓ 1784ms | ✓ 1332ms | ✓ 1070ms | http |
| 38.34.179.20:8445 | ✓ 906ms | 否 | ✓ 307ms | ✓ 883ms | ✓ 760ms | http |
| 38.145.220.11:8445 | 否 | 否 | ✓ 530ms | ✓ 1568ms | ✓ 1441ms | http |
| 219.117.204.211:7799 | 否 | ✓ 1723ms | 否 | ✓ 1077ms | ✓ 939ms | http |
| 38.34.178.154:8445 | ✓ 1521ms | 否 | ✓ 1445ms | ✓ 1959ms | 否 | http |
| 211.217.231.234:8080 | ✓ 809ms | ✓ 1303ms | ✓ 1037ms | ✓ 1108ms | ✓ 871ms | http |
| 183.249.5.117:22222 | 否 | 否 | ✓ 864ms | ✓ 1172ms | ✓ 845ms | http |
| 120.92.212.16:8890 | ✓ 1128ms | ✓ 1372ms | ✓ 1122ms | 否 | ✓ 1305ms | http |
| 120.92.212.16:7890 | ✓ 1087ms | 否 | ✓ 1065ms | 否 | ✓ 1107ms | http |
| 167.103.31.122:8800 | ✓ 1329ms | 否 | ✓ 1703ms | ✓ 1569ms | ✓ 1904ms | http |
| 222.184.48.251:22222 | 否 | ✓ 1323ms | ✓ 1048ms | ✓ 1341ms | 否 | http |
| 38.145.218.229:8450 | 否 | ✓ 1769ms | 否 | ✓ 911ms | ✓ 909ms | http |
| 172.212.68.37:3128 | ✓ 420ms | 否 | ✓ 1614ms | ✓ 1941ms | ✓ 1071ms | http |
| 91.238.105.64:2024 | ✓ 1110ms | 否 | ✓ 1627ms | 否 | ✓ 1533ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1233ms | 否 | ✓ 1294ms | ✓ 1779ms | http |
| 38.34.179.16:8451 | 否 | ✓ 1313ms | ✓ 1703ms | ✓ 1793ms | ✓ 1512ms | http |
| 45.88.0.111:3128 | ✓ 1543ms | ✓ 1537ms | ✓ 1472ms | 否 | 否 | http |
| 38.145.208.184:8452 | ✓ 251ms | ✓ 905ms | ✓ 258ms | ✓ 844ms | ✓ 679ms | http |
| 38.145.208.179:8452 | ✓ 269ms | ✓ 912ms | ✓ 260ms | ✓ 855ms | ✓ 641ms | http |
| 38.145.220.41:8452 | ✓ 290ms | ✓ 842ms | ✓ 247ms | ✓ 1036ms | ✓ 702ms | http |
| 38.145.218.210:8444 | ✓ 333ms | ✓ 799ms | ✓ 410ms | ✓ 876ms | ✓ 719ms | http |
| 38.34.179.167:8450 | ✓ 231ms | ✓ 912ms | ✓ 268ms | ✓ 982ms | ✓ 830ms | http |
| 38.34.179.83:8449 | ✓ 289ms | ✓ 872ms | ✓ 424ms | ✓ 871ms | ✓ 692ms | http |
| 45.136.130.168:8452 | ✓ 257ms | ✓ 901ms | ✓ 246ms | ✓ 927ms | ✓ 697ms | http |
| 45.136.130.174:8450 | ✓ 234ms | ✓ 1312ms | ✓ 261ms | ✓ 857ms | ✓ 701ms | http |
| 38.34.179.181:8446 | ✓ 260ms | ✓ 1407ms | ✓ 346ms | ✓ 842ms | ✓ 669ms | http |
| 38.34.179.48:8448 | ✓ 376ms | ✓ 1048ms | ✓ 290ms | ✓ 1059ms | ✓ 662ms | http |
| 38.145.208.194:8453 | ✓ 305ms | ✓ 1457ms | ✓ 244ms | ✓ 925ms | ✓ 653ms | http |
| 38.145.208.204:8446 | ✓ 456ms | ✓ 904ms | ✓ 746ms | ✓ 885ms | ✓ 664ms | http |
| 38.145.203.32:8452 | ✓ 291ms | ✓ 1827ms | ✓ 241ms | ✓ 856ms | ✓ 653ms | http |
| 38.145.220.81:8446 | ✓ 323ms | ✓ 1889ms | ✓ 267ms | ✓ 892ms | ✓ 663ms | http |
| 45.136.130.171:8452 | ✓ 229ms | ✓ 1936ms | ✓ 248ms | ✓ 864ms | ✓ 758ms | http |
| 38.34.179.25:8448 | ✓ 330ms | 否 | ✓ 254ms | ✓ 884ms | ✓ 672ms | http |
| 38.34.179.21:8452 | ✓ 330ms | 否 | ✓ 272ms | ✓ 912ms | ✓ 678ms | http |
| 38.145.220.33:8448 | ✓ 528ms | ✓ 856ms | ✓ 346ms | ✓ 860ms | ✓ 944ms | http |
| 38.145.208.171:8449 | ✓ 419ms | ✓ 1646ms | ✓ 983ms | ✓ 875ms | ✓ 667ms | http |
| 38.145.220.93:8450 | ✓ 377ms | 否 | ✓ 785ms | ✓ 896ms | ✓ 659ms | http |
| 38.34.179.99:8448 | ✓ 1883ms | ✓ 1333ms | ✓ 255ms | ✓ 890ms | ✓ 713ms | http |
| 38.145.218.101:8448 | ✓ 1111ms | ✓ 920ms | ✓ 339ms | ✓ 1237ms | ✓ 1240ms | http |
| 38.145.218.228:8447 | ✓ 1408ms | 否 | ✓ 1031ms | ✓ 1279ms | ✓ 740ms | http |
| 38.34.179.213:8452 | ✓ 1411ms | 否 | ✓ 1000ms | ✓ 938ms | ✓ 828ms | http |
| 38.34.179.98:8451 | 否 | 否 | ✓ 1426ms | ✓ 946ms | ✓ 673ms | http |
| 38.145.220.56:8453 | 否 | 否 | ✓ 1619ms | ✓ 1921ms | ✓ 1269ms | http |
| 137.220.150.104:6005 | 否 | 否 | ✓ 1468ms | ✓ 1323ms | ✓ 931ms | http |
| 5.102.109.41:999 | 否 | ✓ 1167ms | ✓ 1454ms | 否 | ✓ 1071ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
