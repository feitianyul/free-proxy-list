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

最后更新：2026-05-15 11:21:37 UTC（2026-05-15 19:21:37 UTC+8）

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
| 212.224.88.212:443 | ✓ 762ms | 否 | ✓ 961ms | 否 | ✓ 1581ms | http |
| 1.231.81.166:3128 | ✓ 1588ms | ✓ 1437ms | ✓ 1597ms | ✓ 1178ms | ✓ 901ms | http |
| 129.80.217.21:444 | ✓ 547ms | ✓ 969ms | ✓ 1102ms | ✓ 1217ms | ✓ 955ms | http |
| 119.28.51.157:3128 | ✓ 1201ms | ✓ 1233ms | ✓ 1147ms | ✓ 1157ms | ✓ 1249ms | http |
| 45.59.122.132:80 | ✓ 1162ms | 否 | ✓ 1268ms | 否 | ✓ 1263ms | http |
| 45.88.0.115:3128 | ✓ 568ms | ✓ 1854ms | ✓ 1349ms | 否 | ✓ 1141ms | http |
| 45.125.67.37:8443 | ✓ 1539ms | 否 | ✓ 1335ms | 否 | ✓ 1807ms | http |
| 113.160.132.26:8080 | ✓ 988ms | 否 | 否 | ✓ 1471ms | ✓ 1092ms | http |
| 212.58.132.5:8888 | ✓ 1245ms | 否 | ✓ 1941ms | 否 | ✓ 1893ms | http |
| 103.147.152.12:1080 | ✓ 1187ms | 否 | ✓ 807ms | 否 | ✓ 1594ms | http |
| 113.182.151.78:5108 | ✓ 1373ms | 否 | ✓ 1101ms | ✓ 1790ms | ✓ 1394ms | http |
| 158.160.215.167:8124 | ✓ 761ms | 否 | ✓ 857ms | 否 | ✓ 1633ms | http |
| 181.78.17.131:999 | ✓ 991ms | ✓ 1445ms | ✓ 906ms | ✓ 1639ms | ✓ 1400ms | http |
| 38.188.247.12:999 | 否 | ✓ 1822ms | ✓ 1556ms | ✓ 1872ms | ✓ 1514ms | http |
| 115.231.181.40:8128 | ✓ 1843ms | 否 | 否 | ✓ 1309ms | ✓ 1357ms | http |
| 179.43.159.98:1080 | ✓ 713ms | ✓ 1601ms | ✓ 1170ms | 否 | 否 | http |
| 104.247.51.76:3128 | 否 | 否 | ✓ 993ms | ✓ 1210ms | ✓ 1223ms | http |
| 34.71.229.255:3128 | ✓ 423ms | ✓ 1927ms | ✓ 1097ms | ✓ 1184ms | ✓ 958ms | http |
| 128.199.113.85:9090 | ✓ 1534ms | 否 | 否 | ✓ 1201ms | ✓ 894ms | http |
| 128.199.116.219:9090 | ✓ 1733ms | 否 | ✓ 984ms | ✓ 1563ms | ✓ 1013ms | http |
| 5.75.139.30:1081 | ✓ 1430ms | 否 | ✓ 1543ms | 否 | ✓ 1511ms | http |
| 14.143.222.113:10155 | ✓ 1682ms | 否 | ✓ 1109ms | ✓ 1861ms | 否 | http |
| 128.199.114.189:9090 | 否 | 否 | ✓ 867ms | ✓ 1241ms | ✓ 913ms | http |
| 128.199.254.13:9090 | ✓ 836ms | 否 | ✓ 814ms | ✓ 1170ms | ✓ 940ms | http |
| 146.190.80.158:9090 | ✓ 803ms | 否 | ✓ 829ms | ✓ 1265ms | ✓ 952ms | http |
| 185.191.236.162:3128 | ✓ 1203ms | 否 | ✓ 732ms | ✓ 1719ms | ✓ 1286ms | http |
| 8.154.21.175:3128 | ✓ 927ms | ✓ 1160ms | ✓ 1016ms | ✓ 1236ms | ✓ 975ms | http |
| 157.0.142.246:10057 | ✓ 1153ms | ✓ 1387ms | ✓ 1210ms | 否 | 否 | http |
| 59.46.216.131:30001 | ✓ 1218ms | 否 | ✓ 1105ms | ✓ 1558ms | 否 | http |
| 129.212.224.122:3128 | ✓ 807ms | 否 | 否 | ✓ 1170ms | ✓ 905ms | http |
| 210.223.44.230:3128 | ✓ 1468ms | ✓ 1248ms | ✓ 1812ms | ✓ 1742ms | 否 | http |
| 218.108.131.186:17890 | ✓ 919ms | ✓ 1148ms | ✓ 952ms | ✓ 1182ms | ✓ 962ms | http |
| 152.32.132.190:7890 | 否 | ✓ 1924ms | ✓ 1642ms | ✓ 1883ms | ✓ 752ms | http |
| 150.107.140.238:3128 | ✓ 871ms | 否 | ✓ 950ms | ✓ 1192ms | 否 | http |
| 8.219.97.248:80 | ✓ 1356ms | 否 | ✓ 1634ms | ✓ 1974ms | 否 | http |
| 152.70.91.193:40000 | ✓ 1211ms | 否 | 否 | ✓ 1302ms | ✓ 1247ms | http |
| 65.108.203.36:18080 | ✓ 784ms | 否 | ✓ 745ms | ✓ 1663ms | 否 | http |
| 103.21.220.141:3128 | ✓ 1580ms | 否 | ✓ 808ms | ✓ 1026ms | ✓ 865ms | http |
| 178.213.25.221:7890 | ✓ 1059ms | 否 | ✓ 1548ms | 否 | ✓ 1770ms | http |
| 166.88.55.83:7890 | ✓ 740ms | ✓ 1298ms | ✓ 733ms | ✓ 1053ms | ✓ 1390ms | http |
| 159.89.31.62:8080 | ✓ 525ms | 否 | ✓ 1474ms | ✓ 1690ms | ✓ 1554ms | http |
| 121.130.177.28:8888 | ✓ 1980ms | ✓ 1361ms | ✓ 1525ms | 否 | 否 | http |
| 152.42.170.187:9090 | ✓ 1708ms | 否 | ✓ 1873ms | ✓ 1543ms | ✓ 1019ms | http |
| 222.107.27.7:8017 | 否 | 否 | ✓ 1297ms | ✓ 1730ms | ✓ 1933ms | http |
| 5.252.33.13:2025 | ✓ 1531ms | 否 | ✓ 1560ms | 否 | ✓ 1856ms | http |
| 3.101.133.120:80 | ✓ 632ms | ✓ 1383ms | ✓ 1569ms | ✓ 1309ms | ✓ 965ms | http |
| 77.110.119.136:3128 | ✓ 559ms | ✓ 1613ms | ✓ 1174ms | 否 | 否 | http |
| 128.199.121.61:9090 | ✓ 1007ms | 否 | ✓ 1793ms | ✓ 1158ms | ✓ 976ms | http |
| 120.92.212.16:7890 | ✓ 1154ms | 否 | 否 | ✓ 1658ms | ✓ 1081ms | http |
| 120.92.212.16:8890 | ✓ 1087ms | ✓ 1379ms | ✓ 1050ms | ✓ 1294ms | ✓ 1370ms | http |
| 206.206.126.177:2412 | ✓ 1618ms | 否 | ✓ 1655ms | 否 | ✓ 1614ms | http |
| 150.136.153.231:80 | ✓ 637ms | ✓ 1836ms | 否 | ✓ 1055ms | 否 | http |
| 45.129.141.143:3128 | ✓ 1400ms | 否 | ✓ 1938ms | 否 | ✓ 1647ms | http |
| 104.248.151.93:9090 | ✓ 1797ms | 否 | ✓ 998ms | ✓ 1244ms | ✓ 936ms | http |
| 45.236.129.64:3128 | ✓ 889ms | ✓ 1774ms | ✓ 1154ms | 否 | ✓ 1539ms | http |
| 46.30.46.133:3128 | ✓ 1179ms | 否 | 否 | ✓ 1561ms | ✓ 1527ms | http |
| 158.160.215.167:8126 | ✓ 1409ms | 否 | ✓ 737ms | 否 | ✓ 1644ms | http |
| 91.242.229.129:8092 | ✓ 1249ms | 否 | 否 | ✓ 1651ms | ✓ 1264ms | http |
| 158.160.215.167:8123 | ✓ 1663ms | 否 | ✓ 1250ms | 否 | ✓ 1753ms | http |
| 61.52.131.172:8443 | ✓ 1047ms | ✓ 1242ms | ✓ 1064ms | 否 | ✓ 1045ms | http |
| 103.172.70.173:8080 | 否 | 否 | ✓ 1910ms | ✓ 1500ms | ✓ 1963ms | http |
| 57.129.144.178:40000 | ✓ 1640ms | 否 | ✓ 1614ms | ✓ 1773ms | ✓ 1727ms | http |

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
