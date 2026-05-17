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

最后更新：2026-05-17 05:00:27 UTC（2026-05-17 13:00:27 UTC+8）

**代理总数：67**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 170.106.136.181:31002 | ✓ 1130ms | 否 | ✓ 889ms | ✓ 1309ms | ✓ 466ms | http |
| 113.160.132.26:8080 | ✓ 1893ms | 否 | ✓ 1023ms | ✓ 1351ms | ✓ 1002ms | http |
| 212.58.132.5:8888 | ✓ 1788ms | 否 | 否 | ✓ 1568ms | ✓ 1258ms | http |
| 45.125.67.37:8443 | 否 | 否 | ✓ 1906ms | ✓ 995ms | ✓ 993ms | http |
| 43.156.90.221:10808 | ✓ 694ms | ✓ 1903ms | ✓ 733ms | ✓ 967ms | ✓ 749ms | http |
| 129.80.217.21:444 | 否 | ✓ 1179ms | ✓ 1445ms | ✓ 1241ms | ✓ 953ms | http |
| 34.101.184.164:3128 | ✓ 1489ms | 否 | ✓ 925ms | 否 | ✓ 952ms | http |
| 57.129.144.178:40000 | ✓ 1335ms | 否 | ✓ 1403ms | 否 | ✓ 1931ms | http |
| 129.80.238.83:444 | 否 | ✓ 1183ms | ✓ 1453ms | 否 | ✓ 977ms | http |
| 158.160.215.167:8123 | ✓ 1407ms | ✓ 1947ms | ✓ 1874ms | 否 | 否 | http |
| 217.76.245.80:999 | ✓ 1063ms | 否 | ✓ 1427ms | ✓ 1645ms | ✓ 1627ms | http |
| 65.109.125.111:8443 | ✓ 1349ms | 否 | ✓ 1885ms | 否 | ✓ 1890ms | http |
| 120.92.212.16:8890 | ✓ 1264ms | 否 | ✓ 1599ms | ✓ 1989ms | ✓ 1956ms | http |
| 120.92.212.16:7890 | ✓ 1528ms | ✓ 1157ms | ✓ 1129ms | 否 | 否 | http |
| 115.231.181.40:8128 | 否 | ✓ 1197ms | 否 | ✓ 1319ms | ✓ 997ms | http |
| 91.242.229.129:8092 | ✓ 1131ms | 否 | ✓ 1760ms | 否 | ✓ 1536ms | http |
| 8.154.21.175:3128 | ✓ 798ms | ✓ 988ms | ✓ 1304ms | 否 | ✓ 873ms | http |
| 148.230.4.241:999 | ✓ 661ms | ✓ 1740ms | ✓ 701ms | ✓ 1556ms | ✓ 1298ms | http |
| 8.219.97.248:80 | ✓ 1075ms | 否 | ✓ 1438ms | ✓ 1127ms | 否 | http |
| 84.47.150.125:1080 | ✓ 1119ms | 否 | ✓ 1479ms | 否 | ✓ 1715ms | http |
| 175.194.173.105:3128 | ✓ 1495ms | ✓ 1079ms | ✓ 1644ms | 否 | 否 | http |
| 34.96.238.40:8080 | ✓ 1809ms | ✓ 1157ms | 否 | ✓ 989ms | ✓ 1415ms | http |
| 218.108.131.186:17890 | ✓ 776ms | ✓ 1023ms | ✓ 797ms | ✓ 1012ms | ✓ 845ms | http |
| 128.199.254.13:9090 | 否 | 否 | ✓ 803ms | ✓ 1290ms | ✓ 852ms | http |
| 114.214.165.78:10810 | ✓ 1378ms | ✓ 1657ms | ✓ 1464ms | ✓ 1393ms | ✓ 1424ms | http |
| 103.21.220.141:3128 | ✓ 1750ms | 否 | ✓ 1565ms | ✓ 1027ms | 否 | http |
| 185.40.77.94:1080 | ✓ 1027ms | 否 | ✓ 1953ms | 否 | ✓ 1985ms | http |
| 59.46.216.131:30001 | ✓ 945ms | 否 | ✓ 1047ms | ✓ 1627ms | 否 | http |
| 114.214.170.41:27890 | 否 | ✓ 1284ms | ✓ 1190ms | ✓ 1307ms | ✓ 1063ms | http |
| 168.110.52.228:3128 | ✓ 1503ms | 否 | 否 | ✓ 743ms | ✓ 595ms | http |
| 210.223.44.230:3128 | ✓ 1156ms | ✓ 1111ms | ✓ 1017ms | ✓ 963ms | 否 | http |
| 121.230.9.225:1080 | ✓ 1015ms | 否 | ✓ 1121ms | 否 | ✓ 1882ms | http |
| 101.32.244.83:8080 | ✓ 890ms | ✓ 1594ms | ✓ 881ms | ✓ 1199ms | ✓ 1284ms | http |
| 121.43.196.213:8222 | ✓ 852ms | ✓ 1041ms | ✓ 829ms | ✓ 1081ms | ✓ 874ms | http |
| 121.43.196.210:8222 | ✓ 1980ms | ✓ 991ms | ✓ 1904ms | ✓ 1065ms | ✓ 916ms | http |
| 103.147.152.12:1080 | ✓ 1147ms | ✓ 1858ms | ✓ 1241ms | ✓ 1850ms | 否 | http |
| 77.110.107.80:1080 | ✓ 1939ms | ✓ 1995ms | ✓ 1158ms | 否 | 否 | http |
| 31.172.78.12:3128 | ✓ 1877ms | 否 | ✓ 1188ms | ✓ 1984ms | ✓ 1771ms | http |
| 213.220.3.234:20573 | ✓ 1619ms | ✓ 1610ms | ✓ 1552ms | 否 | ✓ 1721ms | http |
| 152.70.91.193:40000 | 否 | ✓ 1961ms | 否 | ✓ 1740ms | ✓ 1710ms | http |
| 103.147.152.12:1095 | 否 | ✓ 1759ms | ✓ 917ms | ✓ 1980ms | 否 | http |
| 159.223.41.216:9090 | ✓ 709ms | 否 | ✓ 1127ms | ✓ 1030ms | ✓ 806ms | http |
| 152.42.177.32:8888 | ✓ 926ms | 否 | ✓ 1413ms | ✓ 1224ms | 否 | http |
| 47.84.131.156:8100 | 否 | 否 | ✓ 698ms | ✓ 1274ms | ✓ 1005ms | http |
| 64.188.77.221:3128 | 否 | ✓ 1540ms | ✓ 833ms | 否 | ✓ 1856ms | http |
| 3.101.133.120:80 | ✓ 971ms | 否 | ✓ 1832ms | ✓ 1047ms | ✓ 934ms | http |
| 103.220.23.139:8080 | ✓ 1312ms | 否 | ✓ 1265ms | ✓ 1334ms | ✓ 1437ms | http |
| 202.40.186.222:27202 | ✓ 1521ms | 否 | ✓ 1560ms | ✓ 1868ms | ✓ 1447ms | http |
| 116.171.106.26:3443 | ✓ 1398ms | ✓ 1402ms | 否 | 否 | ✓ 1399ms | http |
| 185.21.15.206:3128 | ✓ 1088ms | 否 | ✓ 973ms | 否 | ✓ 1865ms | http |
| 86.104.72.220:1082 | ✓ 1397ms | ✓ 1728ms | ✓ 1491ms | ✓ 1945ms | ✓ 1451ms | http |
| 86.104.72.220:1081 | ✓ 1417ms | ✓ 1735ms | ✓ 1501ms | 否 | ✓ 1349ms | http |
| 91.233.223.147:3128 | ✓ 1187ms | 否 | ✓ 1278ms | 否 | ✓ 1695ms | http |
| 152.32.132.190:7890 | ✓ 1247ms | 否 | ✓ 927ms | 否 | ✓ 1639ms | http |
| 137.59.47.73:3128 | ✓ 1495ms | 否 | 否 | ✓ 1096ms | ✓ 894ms | http |
| 1.231.81.166:3128 | ✓ 1219ms | ✓ 1623ms | 否 | ✓ 1420ms | ✓ 1372ms | http |
| 152.42.170.187:9090 | ✓ 1699ms | 否 | ✓ 1570ms | ✓ 1667ms | 否 | http |
| 5.102.109.41:999 | 否 | 否 | ✓ 1409ms | ✓ 1453ms | ✓ 1391ms | http |
| 2.27.32.81:3128 | ✓ 1689ms | 否 | ✓ 1406ms | 否 | ✓ 1845ms | http |
| 185.200.188.234:10001 | ✓ 1940ms | 否 | ✓ 1491ms | 否 | ✓ 1980ms | http |
| 45.88.0.113:3128 | ✓ 1832ms | 否 | ✓ 1018ms | 否 | ✓ 1287ms | http |
| 116.171.106.111:3443 | 否 | 否 | ✓ 1359ms | ✓ 1589ms | ✓ 1889ms | http |
| 103.129.200.2:8124 | 否 | 否 | ✓ 1404ms | ✓ 1909ms | ✓ 1760ms | http |
| 51.161.50.166:3128 | ✓ 955ms | 否 | 否 | ✓ 1563ms | ✓ 1388ms | http |
| 61.52.131.172:8443 | ✓ 871ms | ✓ 1110ms | ✓ 912ms | 否 | 否 | http |
| 64.181.240.152:3128 | ✓ 331ms | ✓ 796ms | ✓ 296ms | 否 | 否 | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1963ms | ✓ 1717ms | ✓ 1698ms | http |

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
