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

最后更新：2026-03-10 20:50:06 UTC（2026-03-11 04:50:06 UTC+8）

**代理总数：62**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 224ms | ✓ 968ms | ✓ 961ms | ✓ 1239ms | ✓ 878ms | http |
| 45.136.131.47:8443 | 否 | ✓ 833ms | ✓ 927ms | ✓ 974ms | ✓ 671ms | http |
| 35.225.22.61:80 | ✓ 275ms | ✓ 1045ms | ✓ 228ms | ✓ 1174ms | ✓ 1013ms | http |
| 45.140.147.82:1081 | ✓ 915ms | 否 | ✓ 1340ms | 否 | ✓ 1172ms | http |
| 120.92.212.16:7890 | ✓ 1092ms | ✓ 1399ms | ✓ 1084ms | 否 | ✓ 1813ms | http |
| 165.227.5.10:8888 | ✓ 1303ms | ✓ 813ms | ✓ 1210ms | 否 | 否 | http |
| 45.136.131.63:8443 | ✓ 873ms | ✓ 1006ms | ✓ 246ms | ✓ 852ms | ✓ 695ms | http |
| 45.136.130.175:8443 | ✓ 875ms | ✓ 1611ms | ✓ 877ms | ✓ 962ms | ✓ 670ms | http |
| 152.70.98.46:8888 | ✓ 756ms | ✓ 1251ms | ✓ 1519ms | 否 | ✓ 842ms | http |
| 190.9.109.198:999 | ✓ 922ms | ✓ 1352ms | ✓ 1148ms | ✓ 1652ms | ✓ 1486ms | http |
| 81.70.169.194:80 | ✓ 1117ms | ✓ 1391ms | ✓ 1227ms | ✓ 1536ms | ✓ 1129ms | http |
| 101.43.255.96:80 | ✓ 1180ms | ✓ 1476ms | 否 | ✓ 1483ms | ✓ 1172ms | http |
| 45.136.130.191:8443 | 否 | ✓ 776ms | ✓ 720ms | ✓ 930ms | ✓ 726ms | http |
| 45.136.130.188:8443 | 否 | ✓ 797ms | ✓ 694ms | ✓ 901ms | ✓ 781ms | http |
| 202.155.12.161:443 | 否 | 否 | ✓ 1363ms | ✓ 1233ms | ✓ 1054ms | http |
| 190.212.131.238:3128 | 否 | 否 | ✓ 1036ms | ✓ 1925ms | ✓ 1390ms | http |
| 45.186.6.104:3128 | ✓ 1145ms | ✓ 1801ms | ✓ 1859ms | 否 | 否 | http |
| 185.115.74.185:8080 | ✓ 1343ms | ✓ 1932ms | ✓ 1765ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 1359ms | ✓ 1364ms | 否 | ✓ 1413ms | 否 | http |
| 138.124.53.25:7443 | 否 | 否 | ✓ 1319ms | ✓ 1943ms | ✓ 1222ms | http |
| 46.183.25.8:443 | ✓ 1265ms | 否 | ✓ 992ms | ✓ 1076ms | 否 | http |
| 194.213.18.200:443 | ✓ 762ms | 否 | ✓ 1394ms | ✓ 1270ms | ✓ 819ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1303ms | ✓ 984ms | ✓ 1336ms | 否 | http |
| 193.148.58.165:3128 | ✓ 1217ms | 否 | ✓ 1721ms | 否 | ✓ 1334ms | http |
| 121.230.8.181:1080 | ✓ 1264ms | ✓ 1610ms | ✓ 1204ms | ✓ 1573ms | ✓ 1160ms | http |
| 47.77.193.180:1080 | 否 | ✓ 989ms | ✓ 527ms | ✓ 1156ms | ✓ 657ms | http |
| 1.231.81.166:3128 | ✓ 830ms | ✓ 1041ms | ✓ 812ms | ✓ 1069ms | ✓ 791ms | http |
| 162.240.154.26:3128 | ✓ 1636ms | ✓ 1331ms | ✓ 1521ms | ✓ 1575ms | ✓ 1293ms | http |
| 62.113.119.14:8080 | ✓ 1045ms | 否 | ✓ 987ms | ✓ 1558ms | ✓ 1193ms | http |
| 37.139.33.145:1080 | ✓ 1013ms | 否 | ✓ 1395ms | 否 | ✓ 1817ms | http |
| 101.47.73.135:3128 | ✓ 1105ms | 否 | ✓ 1710ms | ✓ 1691ms | 否 | http |
| 158.69.185.37:3129 | ✓ 770ms | 否 | ✓ 752ms | ✓ 1231ms | ✓ 1199ms | http |
| 91.107.141.42:8081 | ✓ 867ms | 否 | ✓ 1352ms | ✓ 1614ms | 否 | http |
| 45.136.130.223:8443 | ✓ 457ms | ✓ 786ms | ✓ 254ms | ✓ 848ms | ✓ 680ms | http |
| 152.42.213.210:8080 | ✓ 1711ms | 否 | ✓ 1108ms | ✓ 1260ms | 否 | http |
| 14.29.168.215:1080 | ✓ 1051ms | ✓ 1589ms | ✓ 1394ms | 否 | 否 | http |
| 121.230.8.144:1080 | ✓ 1668ms | ✓ 1872ms | ✓ 1389ms | ✓ 1725ms | ✓ 1336ms | http |
| 160.19.18.241:3125 | ✓ 1765ms | 否 | 否 | ✓ 1721ms | ✓ 1772ms | http |
| 57.128.188.167:9249 | ✓ 1586ms | 否 | 否 | ✓ 1826ms | ✓ 1683ms | http |
| 172.212.68.37:3128 | ✓ 360ms | ✓ 1628ms | ✓ 1568ms | ✓ 1310ms | ✓ 1319ms | http |
| 39.104.201.40:7890 | ✓ 1809ms | ✓ 1311ms | ✓ 1039ms | 否 | ✓ 1064ms | http |
| 45.136.198.40:3128 | ✓ 951ms | ✓ 1841ms | ✓ 1786ms | 否 | ✓ 1637ms | http |
| 94.176.3.43:7443 | ✓ 1011ms | ✓ 1973ms | ✓ 1710ms | 否 | 否 | http |
| 210.223.44.230:3128 | ✓ 1477ms | 否 | ✓ 677ms | ✓ 1107ms | 否 | http |
| 83.219.250.8:62920 | ✓ 1320ms | ✓ 1691ms | 否 | 否 | ✓ 1721ms | http |
| 177.93.33.55:999 | ✓ 1025ms | 否 | ✓ 1438ms | ✓ 1946ms | 否 | http |
| 178.236.245.17:3128 | ✓ 1499ms | ✓ 1892ms | ✓ 1911ms | 否 | 否 | http |
| 217.77.102.18:3128 | ✓ 1210ms | 否 | ✓ 1618ms | 否 | ✓ 1801ms | http |
| 103.39.51.190:8080 | ✓ 1949ms | 否 | ✓ 1935ms | ✓ 1746ms | ✓ 1490ms | http |
| 89.185.85.138:1080 | ✓ 871ms | ✓ 1449ms | 否 | 否 | ✓ 1208ms | http |
| 178.236.245.59:3128 | ✓ 603ms | ✓ 1790ms | ✓ 1850ms | 否 | 否 | http |
| 67.169.98.211:443 | ✓ 817ms | 否 | ✓ 1128ms | 否 | ✓ 1723ms | http |
| 61.52.131.172:8443 | ✓ 1060ms | ✓ 1296ms | ✓ 1052ms | ✓ 1304ms | ✓ 1037ms | http |
| 8.140.104.98:3128 | 否 | ✓ 1959ms | ✓ 1000ms | ✓ 1374ms | ✓ 1091ms | http |
| 13.36.243.194:9899 | ✓ 1139ms | 否 | 否 | ✓ 1648ms | ✓ 1382ms | http |
| 5.252.33.13:2025 | ✓ 1466ms | 否 | ✓ 1942ms | 否 | ✓ 1792ms | http |
| 168.235.110.63:3128 | ✓ 1219ms | 否 | ✓ 1415ms | ✓ 1882ms | 否 | http |
| 47.105.98.23:3128 | ✓ 1002ms | ✓ 1320ms | ✓ 1103ms | ✓ 1375ms | ✓ 1054ms | http |
| 159.223.42.219:3128 | ✓ 1572ms | 否 | ✓ 1056ms | ✓ 1187ms | ✓ 974ms | http |
| 162.248.165.72:1080 | ✓ 1643ms | 否 | ✓ 1709ms | 否 | ✓ 1152ms | http |
| 111.79.111.126:3128 | ✓ 1166ms | 否 | ✓ 1443ms | ✓ 1423ms | ✓ 1553ms | http |
| 220.121.154.240:3128 | ✓ 1547ms | ✓ 1418ms | 否 | 否 | ✓ 861ms | http |

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
