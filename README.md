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

最后更新：2026-03-14 22:26:48 UTC（2026-03-15 06:26:48 UTC+8）

**代理总数：84**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 83 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 84 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 282ms | ✓ 972ms | ✓ 944ms | ✓ 1339ms | ✓ 854ms | http |
| 38.145.203.135:8443 | ✓ 999ms | ✓ 1599ms | ✓ 1059ms | ✓ 991ms | ✓ 724ms | http |
| 86.53.183.16:1080 | ✓ 1087ms | ✓ 1713ms | ✓ 1795ms | ✓ 1965ms | ✓ 1613ms | http |
| 113.160.132.26:8080 | ✓ 1833ms | ✓ 1510ms | ✓ 1479ms | ✓ 1391ms | ✓ 1079ms | http |
| 45.167.124.52:8080 | 否 | ✓ 1919ms | ✓ 1026ms | 否 | ✓ 1831ms | http |
| 101.47.73.135:3128 | ✓ 1518ms | 否 | ✓ 1104ms | ✓ 1568ms | ✓ 1320ms | http |
| 85.198.96.242:3128 | ✓ 1101ms | 否 | ✓ 1629ms | ✓ 1864ms | 否 | http |
| 165.227.5.10:8888 | ✓ 332ms | ✓ 1360ms | 否 | 否 | ✓ 906ms | http |
| 168.235.110.63:3128 | ✓ 798ms | ✓ 1129ms | ✓ 745ms | ✓ 1022ms | ✓ 785ms | http |
| 101.43.127.100:8877 | ✓ 996ms | ✓ 1249ms | ✓ 1017ms | ✓ 1182ms | ✓ 1013ms | http |
| 81.70.169.194:80 | ✓ 1141ms | ✓ 1539ms | ✓ 1110ms | ✓ 1374ms | ✓ 1163ms | http |
| 150.230.249.50:1080 | 否 | 否 | ✓ 1238ms | ✓ 1141ms | ✓ 936ms | http |
| 38.145.208.138:8447 | ✓ 659ms | ✓ 1057ms | ✓ 954ms | ✓ 1119ms | ✓ 835ms | http |
| 38.145.218.82:8443 | ✓ 630ms | ✓ 1259ms | ✓ 780ms | ✓ 1254ms | ✓ 1059ms | http |
| 116.80.96.108:3172 | ✓ 1640ms | 否 | 否 | ✓ 1994ms | ✓ 1773ms | http |
| 103.84.95.54:7890 | ✓ 1118ms | 否 | ✓ 1331ms | ✓ 1799ms | 否 | http |
| 8.137.112.117:3128 | ✓ 1592ms | ✓ 1454ms | ✓ 1123ms | ✓ 1463ms | ✓ 1156ms | http |
| 62.60.177.204:34094 | ✓ 514ms | 否 | ✓ 816ms | 否 | ✓ 911ms | http |
| 193.23.200.251:10808 | ✓ 1171ms | ✓ 1897ms | ✓ 1758ms | 否 | 否 | http |
| 120.92.212.16:7890 | ✓ 1099ms | 否 | 否 | ✓ 1374ms | ✓ 1386ms | http |
| 121.230.9.136:1080 | ✓ 1322ms | ✓ 1708ms | ✓ 1057ms | 否 | 否 | http |
| 35.225.22.61:80 | ✓ 283ms | ✓ 1116ms | 否 | 否 | ✓ 865ms | http |
| 45.136.130.223:8443 | ✓ 400ms | ✓ 955ms | ✓ 984ms | ✓ 1139ms | ✓ 771ms | http |
| 45.136.130.231:8443 | ✓ 448ms | ✓ 1245ms | ✓ 1117ms | ✓ 1125ms | ✓ 762ms | http |
| 95.3.9.78:8080 | ✓ 1977ms | ✓ 1889ms | ✓ 1932ms | 否 | 否 | http |
| 121.230.9.148:1080 | 否 | 否 | ✓ 1051ms | ✓ 1532ms | ✓ 1125ms | http |
| 103.183.10.169:3125 | ✓ 1568ms | 否 | 否 | ✓ 1611ms | ✓ 1596ms | http |
| 178.236.245.17:3128 | ✓ 1155ms | ✓ 1943ms | ✓ 1038ms | 否 | 否 | http |
| 121.230.9.198:1080 | ✓ 1222ms | ✓ 1522ms | ✓ 1512ms | ✓ 1719ms | ✓ 1277ms | http |
| 121.230.8.111:1080 | ✓ 1472ms | ✓ 1633ms | ✓ 1313ms | ✓ 1832ms | 否 | http |
| 38.145.208.187:8443 | ✓ 965ms | ✓ 927ms | ✓ 854ms | ✓ 992ms | ✓ 758ms | http |
| 88.80.150.82:8080 | ✓ 1050ms | ✓ 1935ms | 否 | 否 | ✓ 1731ms | https |
| 38.145.203.162:8443 | ✓ 787ms | ✓ 1093ms | ✓ 563ms | ✓ 1043ms | ✓ 776ms | http |
| 38.145.208.136:8443 | ✓ 810ms | ✓ 913ms | ✓ 1044ms | ✓ 1066ms | ✓ 963ms | http |
| 45.136.131.28:8447 | ✓ 787ms | ✓ 1172ms | ✓ 1127ms | ✓ 981ms | ✓ 974ms | http |
| 38.145.208.149:8443 | 否 | 否 | ✓ 822ms | ✓ 1039ms | ✓ 811ms | http |
| 38.145.208.99:8443 | ✓ 356ms | ✓ 927ms | ✓ 522ms | ✓ 993ms | ✓ 760ms | http |
| 159.223.42.219:3128 | ✓ 1577ms | 否 | ✓ 1724ms | ✓ 1219ms | ✓ 961ms | http |
| 45.93.28.159:6005 | ✓ 1008ms | 否 | 否 | ✓ 1366ms | ✓ 1883ms | http |
| 45.93.30.177:6005 | ✓ 1624ms | 否 | ✓ 1887ms | ✓ 1946ms | ✓ 1535ms | http |
| 14.225.212.37:7890 | ✓ 932ms | ✓ 1503ms | ✓ 1483ms | ✓ 1473ms | 否 | http |
| 106.117.208.101:7890 | ✓ 1127ms | ✓ 1546ms | ✓ 1480ms | ✓ 1521ms | ✓ 1241ms | http |
| 139.177.190.161:3128 | ✓ 1052ms | 否 | 否 | ✓ 1877ms | ✓ 1609ms | http |
| 24.199.124.152:3128 | ✓ 429ms | ✓ 1037ms | ✓ 1067ms | ✓ 910ms | ✓ 680ms | http |
| 91.247.126.241:2080 | ✓ 1054ms | 否 | ✓ 1266ms | ✓ 1820ms | ✓ 1418ms | http |
| 45.149.92.147:5001 | ✓ 753ms | 否 | ✓ 743ms | ✓ 964ms | ✓ 753ms | http |
| 1.234.153.14:80 | ✓ 1622ms | ✓ 1125ms | ✓ 882ms | ✓ 1158ms | ✓ 803ms | http |
| 139.162.44.152:3128 | ✓ 844ms | 否 | ✓ 872ms | ✓ 1174ms | ✓ 938ms | http |
| 120.92.212.16:8890 | ✓ 1082ms | ✓ 1404ms | ✓ 1157ms | ✓ 1354ms | ✓ 1059ms | http |
| 103.82.93.98:3128 | ✓ 1623ms | 否 | ✓ 956ms | ✓ 1387ms | ✓ 1112ms | http |
| 38.145.208.96:8443 | ✓ 457ms | ✓ 957ms | ✓ 365ms | ✓ 977ms | ✓ 787ms | http |
| 2.56.122.146:10808 | ✓ 514ms | 否 | ✓ 425ms | ✓ 1334ms | ✓ 967ms | http |
| 144.24.29.128:10900 | 否 | ✓ 1568ms | ✓ 1792ms | ✓ 1574ms | 否 | http |
| 95.3.9.78:3128 | ✓ 699ms | ✓ 1820ms | ✓ 1140ms | ✓ 1743ms | ✓ 1442ms | http |
| 38.145.218.101:8447 | ✓ 567ms | ✓ 1881ms | ✓ 348ms | ✓ 1050ms | ✓ 764ms | http |
| 62.113.119.14:8080 | ✓ 948ms | ✓ 1881ms | ✓ 1078ms | 否 | 否 | http |
| 195.158.8.123:3128 | ✓ 1585ms | 否 | ✓ 1673ms | 否 | ✓ 1991ms | http |
| 91.233.223.147:3128 | ✓ 1149ms | ✓ 1963ms | 否 | 否 | ✓ 1913ms | http |
| 38.145.203.43:8443 | 否 | ✓ 1051ms | ✓ 324ms | ✓ 933ms | ✓ 1050ms | http |
| 38.145.203.46:8443 | 否 | ✓ 1013ms | ✓ 308ms | ✓ 947ms | ✓ 1089ms | http |
| 38.145.203.39:8443 | 否 | ✓ 1041ms | ✓ 315ms | ✓ 972ms | ✓ 1043ms | http |
| 38.145.203.41:8443 | 否 | ✓ 1601ms | ✓ 311ms | ✓ 1050ms | ✓ 1887ms | http |
| 89.251.9.11:3128 | ✓ 693ms | ✓ 1131ms | ✓ 926ms | ✓ 1076ms | ✓ 959ms | http |
| 47.101.159.19:8899 | ✓ 1000ms | ✓ 1223ms | ✓ 1038ms | ✓ 1243ms | ✓ 991ms | http |
| 150.249.255.91:3128 | ✓ 1412ms | ✓ 983ms | 否 | ✓ 1869ms | ✓ 841ms | http |
| 1.225.116.115:1080 | ✓ 1788ms | ✓ 1611ms | ✓ 1437ms | ✓ 1579ms | ✓ 1222ms | http |
| 45.136.198.40:3128 | ✓ 1236ms | ✓ 1860ms | ✓ 1905ms | 否 | ✓ 1746ms | http |
| 120.55.163.237:10086 | ✓ 1256ms | 否 | 否 | ✓ 1234ms | ✓ 1779ms | http |
| 103.39.51.190:8080 | ✓ 1904ms | 否 | 否 | ✓ 1671ms | ✓ 1895ms | http |
| 85.208.108.43:2094 | ✓ 651ms | 否 | ✓ 942ms | ✓ 1225ms | ✓ 820ms | http |
| 103.113.70.189:1081 | ✓ 644ms | ✓ 1163ms | 否 | ✓ 1087ms | ✓ 1047ms | http |
| 43.167.227.161:1080 | ✓ 1637ms | 否 | ✓ 598ms | ✓ 929ms | ✓ 733ms | http |
| 138.124.53.25:7443 | 否 | 否 | ✓ 1710ms | ✓ 1778ms | ✓ 1612ms | http |
| 45.136.131.30:8447 | ✓ 988ms | ✓ 893ms | ✓ 421ms | ✓ 1326ms | ✓ 1167ms | http |
| 85.208.108.43:10808 | ✓ 885ms | 否 | ✓ 623ms | ✓ 904ms | ✓ 748ms | http |
| 109.234.38.35:3128 | ✓ 1560ms | ✓ 1817ms | ✓ 1566ms | ✓ 1937ms | ✓ 1307ms | http |
| 172.212.68.37:3128 | ✓ 993ms | ✓ 1462ms | ✓ 839ms | ✓ 1686ms | 否 | http |
| 121.40.231.103:7890 | ✓ 1063ms | ✓ 1260ms | ✓ 1023ms | ✓ 1475ms | ✓ 1002ms | http |
| 61.52.131.172:8443 | ✓ 992ms | ✓ 1318ms | ✓ 1023ms | ✓ 1302ms | ✓ 1018ms | http |
| 38.145.208.144:8443 | ✓ 814ms | ✓ 905ms | ✓ 1060ms | ✓ 988ms | ✓ 862ms | http |
| 38.145.208.148:8443 | ✓ 817ms | ✓ 916ms | ✓ 1047ms | ✓ 983ms | ✓ 884ms | http |
| 38.145.208.37:8443 | ✓ 811ms | ✓ 947ms | ✓ 1021ms | ✓ 995ms | ✓ 898ms | http |
| 45.136.130.228:8443 | ✓ 815ms | ✓ 917ms | ✓ 1047ms | ✓ 979ms | ✓ 923ms | http |
| 64.188.90.36:1080 | ✓ 1505ms | 否 | ✓ 1200ms | ✓ 1970ms | ✓ 1814ms | http |

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
