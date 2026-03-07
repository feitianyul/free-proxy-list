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

最后更新：2026-03-07 23:21:38 UTC（2026-03-08 07:21:38 UTC+8）

**代理总数：91**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 91 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 91 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1687ms | ✓ 947ms | ✓ 1304ms | ✓ 1065ms | ✓ 711ms | http |
| 205.209.118.30:3138 | ✓ 349ms | 否 | ✓ 1249ms | ✓ 1583ms | ✓ 1015ms | http |
| 35.225.22.61:80 | 否 | ✓ 1414ms | ✓ 600ms | ✓ 1088ms | ✓ 866ms | http |
| 1.225.116.115:1080 | ✓ 1142ms | ✓ 1410ms | ✓ 955ms | ✓ 1077ms | ✓ 1111ms | http |
| 202.155.12.161:443 | ✓ 1688ms | ✓ 1946ms | ✓ 1235ms | ✓ 1200ms | 否 | http |
| 211.171.114.154:3128 | ✓ 1954ms | ✓ 1799ms | ✓ 1422ms | 否 | ✓ 1497ms | http |
| 165.227.5.10:8888 | ✓ 1391ms | 否 | ✓ 550ms | 否 | ✓ 690ms | http |
| 117.159.239.49:22222 | ✓ 852ms | ✓ 994ms | ✓ 826ms | ✓ 1054ms | ✓ 853ms | http |
| 103.215.36.88:17304 | ✓ 1013ms | 否 | ✓ 1047ms | ✓ 1201ms | ✓ 1203ms | http |
| 120.92.212.16:7890 | ✓ 960ms | ✓ 1444ms | 否 | 否 | ✓ 904ms | http |
| 103.84.95.54:7890 | ✓ 786ms | 否 | ✓ 1114ms | ✓ 801ms | 否 | http |
| 61.72.221.194:3128 | ✓ 1733ms | ✓ 1991ms | ✓ 1680ms | 否 | ✓ 839ms | http |
| 67.169.98.211:443 | ✓ 1248ms | ✓ 1084ms | 否 | ✓ 1077ms | 否 | http |
| 45.140.147.155:1081 | ✓ 657ms | 否 | ✓ 871ms | 否 | ✓ 1328ms | http |
| 103.145.34.67:8080 | 否 | 否 | ✓ 1167ms | ✓ 1308ms | ✓ 1257ms | http |
| 117.159.239.52:22222 | ✓ 823ms | ✓ 1000ms | ✓ 891ms | ✓ 1039ms | ✓ 809ms | http |
| 101.43.255.96:80 | ✓ 877ms | ✓ 1224ms | ✓ 946ms | ✓ 1146ms | ✓ 896ms | http |
| 81.70.169.194:80 | ✓ 998ms | ✓ 1238ms | ✓ 965ms | ✓ 1250ms | ✓ 894ms | http |
| 121.128.121.54:3128 | ✓ 1567ms | ✓ 1021ms | ✓ 610ms | ✓ 983ms | ✓ 1754ms | http |
| 152.42.213.210:8080 | ✓ 1881ms | 否 | ✓ 1186ms | ✓ 992ms | ✓ 1165ms | http |
| 14.56.107.244:3128 | ✓ 1710ms | 否 | ✓ 908ms | ✓ 1161ms | ✓ 1814ms | http |
| 61.72.221.94:3128 | ✓ 1414ms | 否 | ✓ 1326ms | 否 | ✓ 888ms | http |
| 14.225.217.30:7890 | ✓ 1461ms | 否 | ✓ 891ms | ✓ 1241ms | ✓ 1918ms | http |
| 115.231.181.40:8128 | ✓ 1824ms | ✓ 1283ms | ✓ 1572ms | ✓ 1047ms | ✓ 1257ms | http |
| 113.59.32.148:22222 | ✓ 1089ms | ✓ 1304ms | 否 | 否 | ✓ 1281ms | http |
| 61.72.110.54:3128 | ✓ 1421ms | 否 | 否 | ✓ 1959ms | ✓ 1740ms | http |
| 117.159.239.51:22222 | 否 | ✓ 1086ms | ✓ 841ms | ✓ 1122ms | ✓ 816ms | http |
| 85.9.195.140:1080 | ✓ 1126ms | 否 | ✓ 1786ms | 否 | ✓ 1537ms | http |
| 190.9.109.205:999 | ✓ 961ms | ✓ 1404ms | ✓ 968ms | 否 | 否 | http |
| 116.80.82.229:3172 | ✓ 1459ms | 否 | 否 | ✓ 1809ms | ✓ 1630ms | http |
| 190.9.109.202:999 | ✓ 772ms | ✓ 1335ms | ✓ 1333ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 972ms | ✓ 1737ms | 否 | ✓ 1138ms | 否 | http |
| 125.128.12.14:3128 | ✓ 1302ms | ✓ 1491ms | ✓ 1577ms | ✓ 1133ms | ✓ 783ms | http |
| 125.128.12.144:3128 | ✓ 1103ms | 否 | ✓ 1046ms | 否 | ✓ 956ms | http |
| 59.46.216.131:30001 | ✓ 965ms | ✓ 1229ms | ✓ 948ms | ✓ 1256ms | 否 | http |
| 150.107.140.238:3128 | ✓ 1979ms | 否 | ✓ 1295ms | ✓ 1629ms | 否 | http |
| 138.124.53.25:7443 | ✓ 1431ms | 否 | ✓ 1316ms | ✓ 1621ms | 否 | http |
| 168.235.110.63:3128 | ✓ 1718ms | ✓ 1301ms | ✓ 616ms | ✓ 1296ms | ✓ 956ms | http |
| 14.225.222.164:7890 | ✓ 1968ms | 否 | ✓ 897ms | 否 | ✓ 1480ms | http |
| 194.59.204.87:9080 | ✓ 1106ms | ✓ 1792ms | ✓ 1783ms | 否 | 否 | http |
| 46.183.25.8:443 | ✓ 1206ms | 否 | ✓ 1592ms | ✓ 1499ms | ✓ 1685ms | http |
| 159.223.42.219:3128 | ✓ 896ms | 否 | ✓ 1322ms | ✓ 1025ms | ✓ 825ms | http |
| 116.80.82.228:3172 | ✓ 1459ms | 否 | 否 | ✓ 1873ms | ✓ 1779ms | http |
| 46.249.103.192:443 | ✓ 1060ms | 否 | ✓ 1457ms | ✓ 1780ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1564ms | ✓ 1901ms | ✓ 1224ms | 否 | ✓ 1463ms | http |
| 117.72.46.30:7890 | ✓ 942ms | ✓ 1650ms | ✓ 1013ms | 否 | 否 | http |
| 120.240.29.168:22222 | ✓ 891ms | ✓ 1153ms | ✓ 966ms | 否 | 否 | http |
| 117.159.239.50:22222 | ✓ 783ms | ✓ 1105ms | ✓ 831ms | ✓ 1054ms | ✓ 832ms | http |
| 117.159.239.44:22222 | ✓ 759ms | ✓ 977ms | ✓ 972ms | ✓ 1023ms | ✓ 827ms | http |
| 120.240.35.177:22222 | ✓ 847ms | ✓ 1199ms | ✓ 955ms | ✓ 1035ms | ✓ 892ms | http |
| 120.240.35.176:22222 | ✓ 868ms | ✓ 1206ms | ✓ 928ms | ✓ 1105ms | ✓ 874ms | http |
| 113.59.32.161:22222 | ✓ 984ms | ✓ 1176ms | ✓ 1017ms | ✓ 1137ms | ✓ 991ms | http |
| 222.184.48.236:22222 | 否 | 否 | ✓ 1280ms | ✓ 1172ms | ✓ 850ms | http |
| 113.59.32.142:22222 | ✓ 977ms | ✓ 1302ms | ✓ 937ms | ✓ 1230ms | ✓ 927ms | http |
| 120.240.29.173:22222 | ✓ 871ms | ✓ 1168ms | ✓ 935ms | ✓ 1087ms | ✓ 883ms | http |
| 180.76.115.231:3128 | ✓ 900ms | ✓ 1207ms | ✓ 1085ms | ✓ 1226ms | ✓ 977ms | http |
| 113.59.32.163:22222 | ✓ 924ms | 否 | ✓ 852ms | ✓ 1200ms | ✓ 898ms | http |
| 183.249.5.111:22222 | 否 | 否 | ✓ 743ms | ✓ 905ms | ✓ 708ms | http |
| 120.198.141.75:22222 | ✓ 959ms | ✓ 1445ms | ✓ 1092ms | ✓ 1333ms | ✓ 1010ms | http |
| 121.230.8.181:1080 | ✓ 962ms | ✓ 1475ms | ✓ 1168ms | ✓ 1496ms | ✓ 1134ms | http |
| 222.184.48.242:22222 | ✓ 1132ms | ✓ 1647ms | ✓ 1946ms | ✓ 1180ms | ✓ 1323ms | http |
| 39.104.201.40:7890 | ✓ 872ms | ✓ 1132ms | ✓ 1839ms | ✓ 1995ms | ✓ 922ms | http |
| 185.243.218.43:49153 | ✓ 850ms | 否 | ✓ 1762ms | 否 | ✓ 1907ms | http |
| 42.115.72.27:2064 | 否 | 否 | ✓ 1793ms | ✓ 1797ms | ✓ 1522ms | http |
| 148.153.56.51:80 | ✓ 634ms | ✓ 585ms | ✓ 1231ms | ✓ 1165ms | ✓ 1821ms | http |
| 94.79.152.14:80 | ✓ 896ms | 否 | ✓ 1663ms | ✓ 1584ms | 否 | http |
| 152.42.213.210:80 | ✓ 1628ms | 否 | ✓ 1248ms | 否 | ✓ 1805ms | http |
| 45.140.147.82:1081 | ✓ 596ms | 否 | ✓ 1779ms | ✓ 1458ms | ✓ 1255ms | http |
| 103.215.36.88:15968 | ✓ 960ms | ✓ 1242ms | ✓ 1002ms | ✓ 1157ms | ✓ 980ms | http |
| 103.215.36.88:17565 | ✓ 880ms | ✓ 1229ms | 否 | ✓ 1201ms | ✓ 961ms | http |
| 51.250.37.15:6666 | ✓ 1708ms | 否 | ✓ 1743ms | 否 | ✓ 1947ms | http |
| 47.77.193.180:1080 | ✓ 527ms | ✓ 753ms | ✓ 140ms | ✓ 670ms | ✓ 487ms | http |
| 103.215.36.88:10864 | 否 | 否 | ✓ 1969ms | ✓ 1907ms | ✓ 1076ms | http |
| 103.215.36.88:16299 | ✓ 934ms | ✓ 1314ms | ✓ 1117ms | ✓ 1394ms | ✓ 960ms | http |
| 103.215.36.88:15247 | ✓ 1143ms | ✓ 1609ms | ✓ 1094ms | 否 | ✓ 1014ms | http |
| 45.140.147.82:1082 | ✓ 1128ms | ✓ 1285ms | 否 | 否 | ✓ 1264ms | http |
| 178.236.245.17:3128 | ✓ 997ms | ✓ 1927ms | ✓ 1978ms | 否 | 否 | http |
| 54.222.174.194:80 | 否 | 否 | ✓ 1583ms | ✓ 1874ms | ✓ 1461ms | http |
| 201.144.20.238:3128 | ✓ 1258ms | ✓ 1269ms | ✓ 1260ms | ✓ 1469ms | ✓ 916ms | http |
| 38.180.2.107:3128 | ✓ 1002ms | ✓ 1928ms | ✓ 1660ms | 否 | ✓ 1839ms | http |
| 193.168.173.136:443 | ✓ 1146ms | 否 | ✓ 1516ms | 否 | ✓ 1815ms | http |
| 114.4.251.26:8080 | ✓ 1882ms | 否 | ✓ 1138ms | ✓ 1338ms | ✓ 1302ms | http |
| 112.78.187.186:8080 | ✓ 1883ms | 否 | ✓ 1282ms | ✓ 1291ms | ✓ 1278ms | http |
| 106.14.203.63:3333 | ✓ 1036ms | 否 | ✓ 1943ms | ✓ 1299ms | ✓ 1061ms | http |
| 103.39.51.190:8080 | ✓ 1374ms | 否 | 否 | ✓ 1320ms | ✓ 1488ms | http |
| 162.240.154.26:3128 | 否 | ✓ 1440ms | ✓ 1568ms | 否 | ✓ 1302ms | http |
| 45.136.198.40:3128 | ✓ 1336ms | 否 | ✓ 1111ms | ✓ 1693ms | ✓ 1390ms | http |
| 116.80.82.224:3172 | 否 | 否 | ✓ 1975ms | ✓ 1872ms | ✓ 1858ms | http |
| 120.240.35.175:22222 | ✓ 979ms | ✓ 1222ms | ✓ 935ms | ✓ 1081ms | ✓ 903ms | http |
| 103.183.10.169:3125 | ✓ 1828ms | 否 | 否 | ✓ 1551ms | ✓ 1750ms | http |
| 103.215.36.88:15879 | ✓ 1259ms | 否 | ✓ 1019ms | 否 | ✓ 894ms | http |

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
